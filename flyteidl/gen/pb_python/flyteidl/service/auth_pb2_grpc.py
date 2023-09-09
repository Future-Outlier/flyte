# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from flyteidl.service import auth_pb2 as flyteidl_dot_service_dot_auth__pb2


class AuthMetadataServiceStub(object):
    """The following defines an RPC service that is also served over HTTP via grpc-gateway.
    Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
    RPCs defined in this service must be anonymously accessible.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetOAuth2Metadata = channel.unary_unary(
                '/flyteidl.service.AuthMetadataService/GetOAuth2Metadata',
                request_serializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataRequest.SerializeToString,
                response_deserializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataResponse.FromString,
                )
        self.GetPublicClientConfig = channel.unary_unary(
                '/flyteidl.service.AuthMetadataService/GetPublicClientConfig',
                request_serializer=flyteidl_dot_service_dot_auth__pb2.PublicClientAuthConfigRequest.SerializeToString,
                response_deserializer=flyteidl_dot_service_dot_auth__pb2.PublicClientAuthConfigResponse.FromString,
                )


class AuthMetadataServiceServicer(object):
    """The following defines an RPC service that is also served over HTTP via grpc-gateway.
    Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
    RPCs defined in this service must be anonymously accessible.
    """

    def GetOAuth2Metadata(self, request, context):
        """Anonymously accessible. Retrieves local or external oauth authorization server metadata.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPublicClientConfig(self, request, context):
        """Anonymously accessible. Retrieves the client information clients should use when initiating OAuth2 authorization
        requests.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthMetadataServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetOAuth2Metadata': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOAuth2Metadata,
                    request_deserializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataRequest.FromString,
                    response_serializer=flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataResponse.SerializeToString,
            ),
            'GetPublicClientConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPublicClientConfig,
                    request_deserializer=flyteidl_dot_service_dot_auth__pb2.PublicClientAuthConfigRequest.FromString,
                    response_serializer=flyteidl_dot_service_dot_auth__pb2.PublicClientAuthConfigResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'flyteidl.service.AuthMetadataService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AuthMetadataService(object):
    """The following defines an RPC service that is also served over HTTP via grpc-gateway.
    Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
    RPCs defined in this service must be anonymously accessible.
    """

    @staticmethod
    def GetOAuth2Metadata(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/flyteidl.service.AuthMetadataService/GetOAuth2Metadata',
            flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataRequest.SerializeToString,
            flyteidl_dot_service_dot_auth__pb2.OAuth2MetadataResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPublicClientConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/flyteidl.service.AuthMetadataService/GetPublicClientConfig',
            flyteidl_dot_service_dot_auth__pb2.PublicClientAuthConfigRequest.SerializeToString,
            flyteidl_dot_service_dot_auth__pb2.PublicClientAuthConfigResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
