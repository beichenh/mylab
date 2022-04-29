# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: server.proto for package 'gitaly'

require 'grpc'
require 'server_pb'

module Gitaly
  module ServerService
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.ServerService'

      rpc :ServerInfo, ::Gitaly::ServerInfoRequest, ::Gitaly::ServerInfoResponse
      rpc :DiskStatistics, ::Gitaly::DiskStatisticsRequest, ::Gitaly::DiskStatisticsResponse
      # ClockSynced checks if machine clock is synced
      # (the offset is less that the one passed in the request).
      rpc :ClockSynced, ::Gitaly::ClockSyncedRequest, ::Gitaly::ClockSyncedResponse
    end

    Stub = Service.rpc_stub_class
  end
end