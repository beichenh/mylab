//go:build static && system_libgit2
// +build static,system_libgit2

package main

import (
	"fmt"
	"testing"
	"time"

	git "github.com/libgit2/git2go/v33"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/gitlab-org/gitaly/v14/cmd/gitaly-git2go-v14/git2goutil"
	cmdtesthelper "gitlab.com/gitlab-org/gitaly/v14/cmd/gitaly-git2go-v14/testhelper"
	"gitlab.com/gitlab-org/gitaly/v14/internal/git/gittest"
	"gitlab.com/gitlab-org/gitaly/v14/internal/git2go"
	"gitlab.com/gitlab-org/gitaly/v14/internal/testhelper"
	"gitlab.com/gitlab-org/gitaly/v14/internal/testhelper/testcfg"
)

func TestMerge_missingArguments(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	cfg, repo, repoPath := testcfg.BuildWithRepo(t)
	executor := buildExecutor(t, cfg)

	testcases := []struct {
		desc        string
		request     git2go.MergeCommand
		expectedErr string
	}{
		{
			desc:        "no arguments",
			expectedErr: "merge: invalid parameters: missing repository",
		},
		{
			desc:        "missing repository",
			request:     git2go.MergeCommand{AuthorName: "Foo", AuthorMail: "foo@example.com", Message: "Foo", Ours: "HEAD", Theirs: "HEAD"},
			expectedErr: "merge: invalid parameters: missing repository",
		},
		{
			desc:        "missing author name",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorMail: "foo@example.com", Message: "Foo", Ours: "HEAD", Theirs: "HEAD"},
			expectedErr: "merge: invalid parameters: missing author name",
		},
		{
			desc:        "missing author mail",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorName: "Foo", Message: "Foo", Ours: "HEAD", Theirs: "HEAD"},
			expectedErr: "merge: invalid parameters: missing author mail",
		},
		{
			desc:        "missing message",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorName: "Foo", AuthorMail: "foo@example.com", Ours: "HEAD", Theirs: "HEAD"},
			expectedErr: "merge: invalid parameters: missing message",
		},
		{
			desc:        "missing ours",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorName: "Foo", AuthorMail: "foo@example.com", Message: "Foo", Theirs: "HEAD"},
			expectedErr: "merge: invalid parameters: missing ours",
		},
		{
			desc:        "missing theirs",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorName: "Foo", AuthorMail: "foo@example.com", Message: "Foo", Ours: "HEAD"},
			expectedErr: "merge: invalid parameters: missing theirs",
		},
		// Committer* arguments are required only when at least one of them is non-empty
		{
			desc:        "missing committer mail",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorName: "Foo", AuthorMail: "foo@example.com", CommitterName: "Bar", Message: "Foo", Theirs: "HEAD", Ours: "HEAD"},
			expectedErr: "merge: invalid parameters: missing committer mail",
		},
		{
			desc:        "missing committer name",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorName: "Foo", AuthorMail: "foo@example.com", CommitterMail: "bar@example.com", Message: "Foo", Theirs: "HEAD", Ours: "HEAD"},
			expectedErr: "merge: invalid parameters: missing committer name",
		},
		{
			desc:        "missing committer date",
			request:     git2go.MergeCommand{Repository: repoPath, AuthorName: "Foo", AuthorMail: "foo@example.com", CommitterName: "Bar", CommitterMail: "bar@example.com", Message: "Foo", Theirs: "HEAD", Ours: "HEAD"},
			expectedErr: "merge: invalid parameters: missing committer date",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := executor.Merge(ctx, repo, tc.request)
			require.Error(t, err)
			require.Equal(t, tc.expectedErr, err.Error())
		})
	}
}

func TestMerge_invalidRepositoryPath(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	cfg, repo, _ := testcfg.BuildWithRepo(t)
	testcfg.BuildGitalyGit2Go(t, cfg)
	executor := buildExecutor(t, cfg)

	_, err := executor.Merge(ctx, repo, git2go.MergeCommand{
		Repository: "/does/not/exist", AuthorName: "Foo", AuthorMail: "foo@example.com", Message: "Foo", Ours: "HEAD", Theirs: "HEAD",
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "merge: could not open repository")
}

func TestMerge_trees(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	testcases := []struct {
		desc             string
		base             map[string]string
		ours             map[string]string
		theirs           map[string]string
		expected         map[string]string
		withCommitter    bool
		squash           bool
		expectedResponse git2go.MergeResult
		expectedErr      error
	}{
		{
			desc: "trivial merge succeeds",
			base: map[string]string{
				"file": "a",
			},
			ours: map[string]string{
				"file": "a",
			},
			theirs: map[string]string{
				"file": "a",
			},
			expected: map[string]string{
				"file": "a",
			},
			expectedResponse: git2go.MergeResult{
				CommitID: "7d5ae8fb6d2b301c53560bd728004d77778998df",
			},
		},
		{
			desc: "trivial merge with different committer succeeds",
			base: map[string]string{
				"file": "a",
			},
			ours: map[string]string{
				"file": "a",
			},
			theirs: map[string]string{
				"file": "a",
			},
			expected: map[string]string{
				"file": "a",
			},
			withCommitter: true,
			expectedResponse: git2go.MergeResult{
				CommitID: "cba8c5ddf5a5a24f2f606e4b62d348feb1214b70",
			},
		},
		{
			desc: "trivial squash succeeds",
			base: map[string]string{
				"file": "a",
			},
			ours: map[string]string{
				"file": "a",
			},
			theirs: map[string]string{
				"file": "a",
			},
			expected: map[string]string{
				"file": "a",
			},
			squash: true,
			expectedResponse: git2go.MergeResult{
				CommitID: "d4c52f063cd6544959d6b0d9a3d8fa8463c34086",
			},
		},
		{
			desc: "non-trivial merge succeeds",
			base: map[string]string{
				"file": "a\nb\nc\nd\ne\nf\n",
			},
			ours: map[string]string{
				"file": "0\na\nb\nc\nd\ne\nf\n",
			},
			theirs: map[string]string{
				"file": "a\nb\nc\nd\ne\nf\n0\n",
			},
			expected: map[string]string{
				"file": "0\na\nb\nc\nd\ne\nf\n0\n",
			},
			expectedResponse: git2go.MergeResult{
				CommitID: "348b9b489c3ca128a4555c7a51b20335262519c7",
			},
		},
		{
			desc: "non-trivial squash succeeds",
			base: map[string]string{
				"file": "a\nb\nc\nd\ne\nf\n",
			},
			ours: map[string]string{
				"file": "0\na\nb\nc\nd\ne\nf\n",
			},
			theirs: map[string]string{
				"file": "a\nb\nc\nd\ne\nf\n0\n",
			},
			expected: map[string]string{
				"file": "0\na\nb\nc\nd\ne\nf\n0\n",
			},
			squash: true,
			expectedResponse: git2go.MergeResult{
				CommitID: "7ef7460f69503265a247e06218391cfa57c521fc",
			},
		},
		{
			desc: "multiple files succeed",
			base: map[string]string{
				"1": "foo",
				"2": "bar",
				"3": "qux",
			},
			ours: map[string]string{
				"1": "foo",
				"2": "modified",
				"3": "qux",
			},
			theirs: map[string]string{
				"1": "modified",
				"2": "bar",
				"3": "qux",
			},
			expected: map[string]string{
				"1": "modified",
				"2": "modified",
				"3": "qux",
			},
			expectedResponse: git2go.MergeResult{
				CommitID: "e9be4578f89ea52d44936fb36517e837d698b34b",
			},
		},
		{
			desc: "multiple files squash succeed",
			base: map[string]string{
				"1": "foo",
				"2": "bar",
				"3": "qux",
			},
			ours: map[string]string{
				"1": "foo",
				"2": "modified",
				"3": "qux",
			},
			theirs: map[string]string{
				"1": "modified",
				"2": "bar",
				"3": "qux",
			},
			expected: map[string]string{
				"1": "modified",
				"2": "modified",
				"3": "qux",
			},
			squash: true,
			expectedResponse: git2go.MergeResult{
				CommitID: "a680459fe541be728c8494fb76c233a344c04460",
			},
		},
		{
			desc: "conflicting merge fails",
			base: map[string]string{
				"1": "foo",
			},
			ours: map[string]string{
				"1": "bar",
			},
			theirs: map[string]string{
				"1": "qux",
			},
			expectedErr: fmt.Errorf("merge: %w", git2go.ConflictingFilesError{
				ConflictingFiles: []string{"1"},
			}),
		},
	}

	for _, tc := range testcases {
		cfg, repoProto, repoPath := testcfg.BuildWithRepo(t)
		testcfg.BuildGitalyGit2Go(t, cfg)
		executor := buildExecutor(t, cfg)

		base := cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{nil}, tc.base)
		ours := cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{base}, tc.ours)
		theirs := cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{base}, tc.theirs)

		authorDate := time.Date(2020, 7, 30, 7, 45, 50, 0, time.FixedZone("UTC+2", +2*60*60))
		committerDate := time.Date(2021, 7, 30, 7, 45, 50, 0, time.FixedZone("UTC+2", +2*60*60))

		t.Run(tc.desc, func(t *testing.T) {
			mergeCommand := git2go.MergeCommand{
				Repository: repoPath,
				AuthorName: "John Doe",
				AuthorMail: "john.doe@example.com",
				AuthorDate: authorDate,
				Message:    "Merge message",
				Ours:       ours.String(),
				Theirs:     theirs.String(),
				Squash:     tc.squash,
			}
			if tc.withCommitter {
				mergeCommand.CommitterName = "Jane Doe"
				mergeCommand.CommitterMail = "jane.doe@example.com"
				mergeCommand.CommitterDate = committerDate
			}
			response, err := executor.Merge(ctx, repoProto, mergeCommand)

			if tc.expectedErr != nil {
				require.Equal(t, tc.expectedErr, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResponse, response)

			repo, err := git2goutil.OpenRepository(repoPath)
			require.NoError(t, err)
			defer repo.Free()

			commitOid, err := git.NewOid(response.CommitID)
			require.NoError(t, err)

			commit, err := repo.LookupCommit(commitOid)
			require.NoError(t, err)

			tree, err := commit.Tree()
			require.NoError(t, err)
			require.EqualValues(t, len(tc.expected), tree.EntryCount())

			for name, contents := range tc.expected {
				entry := tree.EntryByName(name)
				require.NotNil(t, entry)

				blob, err := repo.LookupBlob(entry.Id)
				require.NoError(t, err)
				require.Equal(t, []byte(contents), blob.Contents())
			}
		})
	}
}

func TestMerge_squash(t *testing.T) {
	t.Parallel()

	ctx := testhelper.Context(t)

	cfg, repoProto, repoPath := testcfg.BuildWithRepo(t)
	testcfg.BuildGitalyGit2Go(t, cfg)
	executor := buildExecutor(t, cfg)

	baseFiles := map[string]string{"file.txt": "b\nc"}
	ourFiles := map[string]string{"file.txt": "a\nb\nc"}
	theirFiles1 := map[string]string{"file.txt": "b\nc\nd"}
	theirFiles2 := map[string]string{"file.txt": "b\nc\nd\ne"}

	base := cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{nil}, baseFiles)
	ours := cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{base}, ourFiles)
	theirs1 := cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{base}, theirFiles1)
	theirs2 := cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{theirs1}, theirFiles2)

	date := time.Date(2020, 7, 30, 7, 45, 50, 0, time.FixedZone("UTC+2", +2*60*60))
	response, err := executor.Merge(ctx, repoProto, git2go.MergeCommand{
		Repository: repoPath,
		AuthorName: "John Doe",
		AuthorMail: "john.doe@example.com",
		AuthorDate: date,
		Message:    "Merge message",
		Ours:       ours.String(),
		Theirs:     theirs2.String(),
		Squash:     true,
	})
	require.NoError(t, err)
	assert.Equal(t, "027d909803fbb3d17c3b10c1dfe8f120d99392e4", response.CommitID)

	repo, err := git2goutil.OpenRepository(repoPath)
	require.NoError(t, err)

	commitOid, err := git.NewOid(response.CommitID)
	require.NoError(t, err)

	isDescendant, err := repo.DescendantOf(commitOid, theirs2)
	require.NoError(t, err)
	require.False(t, isDescendant)

	commit, err := repo.LookupCommit(commitOid)
	require.NoError(t, err)

	require.Equal(t, uint(1), commit.ParentCount())
	require.Equal(t, ours, commit.ParentId(0))

	tree, err := commit.Tree()
	require.NoError(t, err)

	entry := tree.EntryByName("file.txt")
	require.NotNil(t, entry)

	blob, err := repo.LookupBlob(entry.Id)
	require.NoError(t, err)
	require.Equal(t, "a\nb\nc\nd\ne", string(blob.Contents()))
}

func TestMerge_recursive(t *testing.T) {
	t.Parallel()
	ctx := testhelper.Context(t)

	cfg := testcfg.Build(t)
	testcfg.BuildGitalyGit2Go(t, cfg)
	executor := buildExecutor(t, cfg)

	repoProto, repoPath := gittest.InitRepo(t, cfg, cfg.Storages[0])

	base := cmdtesthelper.BuildCommit(t, repoPath, nil, map[string]string{"base": "base\n"})

	oursContents := map[string]string{"base": "base\n", "ours": "ours-0\n"}
	ours := make([]*git.Oid, git2go.MergeRecursionLimit)
	ours[0] = cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{base}, oursContents)

	theirsContents := map[string]string{"base": "base\n", "theirs": "theirs-0\n"}
	theirs := make([]*git.Oid, git2go.MergeRecursionLimit)
	theirs[0] = cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{base}, theirsContents)

	// We're now creating a set of criss-cross merges which look like the following graph:
	//
	//        o---o---o---o---o-   -o---o ours
	//       / \ / \ / \ / \ / \ . / \ /
	// base o   X   X   X   X    .    X
	//       \ / \ / \ / \ / \ / . \ / \
	//        o---o---o---o---o-   -o---o theirs
	//
	// We then merge ours with theirs. The peculiarity about this merge is that the merge base
	// is not unique, and as a result the merge will generate virtual merge bases for each of
	// the criss-cross merges. This operation may thus be heavily expensive to perform.
	for i := 1; i < git2go.MergeRecursionLimit; i++ {
		oursContents["ours"] = fmt.Sprintf("ours-%d\n", i)
		oursContents["theirs"] = fmt.Sprintf("theirs-%d\n", i-1)
		theirsContents["ours"] = fmt.Sprintf("ours-%d\n", i-1)
		theirsContents["theirs"] = fmt.Sprintf("theirs-%d\n", i)

		ours[i] = cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{ours[i-1], theirs[i-1]}, oursContents)
		theirs[i] = cmdtesthelper.BuildCommit(t, repoPath, []*git.Oid{theirs[i-1], ours[i-1]}, theirsContents)
	}

	authorDate := time.Date(2020, 7, 30, 7, 45, 50, 0, time.FixedZone("UTC+2", +2*60*60))

	// When creating the criss-cross merges, we have been doing evil merges
	// as each merge has applied changes from the other side while at the
	// same time incrementing the own file contents. As we exceed the merge
	// limit, git will just pick one of both possible merge bases when
	// hitting that limit instead of computing another virtual merge base.
	// The result is thus a merge of the following three commits:
	//
	// merge base           ours                theirs
	// ----------           ----                ------
	//
	// base:   "base"       base:   "base"      base:   "base"
	// theirs: "theirs-1"   theirs: "theirs-1   theirs: "theirs-2"
	// ours:   "ours-0"     ours:   "ours-2"    ours:   "ours-1"
	//
	// This is a classical merge commit as "ours" differs in all three
	// cases. We thus expect a merge conflict, which unfortunately
	// demonstrates that restricting the recursion limit may cause us to
	// fail resolution.
	_, err := executor.Merge(ctx, repoProto, git2go.MergeCommand{
		Repository: repoPath,
		AuthorName: "John Doe",
		AuthorMail: "john.doe@example.com",
		AuthorDate: authorDate,
		Message:    "Merge message",
		Ours:       ours[len(ours)-1].String(),
		Theirs:     theirs[len(theirs)-1].String(),
	})
	require.Equal(t, fmt.Errorf("merge: %w", git2go.ConflictingFilesError{
		ConflictingFiles: []string{"theirs"},
	}), err)

	// Otherwise, if we're merging an earlier criss-cross merge which has
	// half of the limit many criss-cross patterns, we exactly hit the
	// recursion limit and thus succeed.
	response, err := executor.Merge(ctx, repoProto, git2go.MergeCommand{
		Repository: repoPath,
		AuthorName: "John Doe",
		AuthorMail: "john.doe@example.com",
		AuthorDate: authorDate,
		Message:    "Merge message",
		Ours:       ours[git2go.MergeRecursionLimit/2].String(),
		Theirs:     theirs[git2go.MergeRecursionLimit/2].String(),
	})
	require.NoError(t, err)

	repo, err := git2goutil.OpenRepository(repoPath)
	require.NoError(t, err)

	commitOid, err := git.NewOid(response.CommitID)
	require.NoError(t, err)

	commit, err := repo.LookupCommit(commitOid)
	require.NoError(t, err)

	tree, err := commit.Tree()
	require.NoError(t, err)

	require.EqualValues(t, 3, tree.EntryCount())
	for name, contents := range map[string]string{
		"base":   "base\n",
		"ours":   "ours-10\n",
		"theirs": "theirs-10\n",
	} {
		entry := tree.EntryByName(name)
		require.NotNil(t, entry)

		blob, err := repo.LookupBlob(entry.Id)
		require.NoError(t, err)
		require.Equal(t, []byte(contents), blob.Contents())
	}
}
