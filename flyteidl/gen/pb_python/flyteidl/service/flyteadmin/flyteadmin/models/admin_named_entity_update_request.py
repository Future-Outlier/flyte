# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.admin_named_entity_identifier import AdminNamedEntityIdentifier  # noqa: F401,E501
from flyteadmin.models.admin_named_entity_metadata import AdminNamedEntityMetadata  # noqa: F401,E501
from flyteadmin.models.core_resource_type import CoreResourceType  # noqa: F401,E501


class AdminNamedEntityUpdateRequest(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'resource_type': 'CoreResourceType',
        'id': 'AdminNamedEntityIdentifier',
        'metadata': 'AdminNamedEntityMetadata'
    }

    attribute_map = {
        'resource_type': 'resource_type',
        'id': 'id',
        'metadata': 'metadata'
    }

    def __init__(self, resource_type=None, id=None, metadata=None):  # noqa: E501
        """AdminNamedEntityUpdateRequest - a model defined in Swagger"""  # noqa: E501

        self._resource_type = None
        self._id = None
        self._metadata = None
        self.discriminator = None

        if resource_type is not None:
            self.resource_type = resource_type
        if id is not None:
            self.id = id
        if metadata is not None:
            self.metadata = metadata

    @property
    def resource_type(self):
        """Gets the resource_type of this AdminNamedEntityUpdateRequest.  # noqa: E501


        :return: The resource_type of this AdminNamedEntityUpdateRequest.  # noqa: E501
        :rtype: CoreResourceType
        """
        return self._resource_type

    @resource_type.setter
    def resource_type(self, resource_type):
        """Sets the resource_type of this AdminNamedEntityUpdateRequest.


        :param resource_type: The resource_type of this AdminNamedEntityUpdateRequest.  # noqa: E501
        :type: CoreResourceType
        """

        self._resource_type = resource_type

    @property
    def id(self):
        """Gets the id of this AdminNamedEntityUpdateRequest.  # noqa: E501


        :return: The id of this AdminNamedEntityUpdateRequest.  # noqa: E501
        :rtype: AdminNamedEntityIdentifier
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this AdminNamedEntityUpdateRequest.


        :param id: The id of this AdminNamedEntityUpdateRequest.  # noqa: E501
        :type: AdminNamedEntityIdentifier
        """

        self._id = id

    @property
    def metadata(self):
        """Gets the metadata of this AdminNamedEntityUpdateRequest.  # noqa: E501


        :return: The metadata of this AdminNamedEntityUpdateRequest.  # noqa: E501
        :rtype: AdminNamedEntityMetadata
        """
        return self._metadata

    @metadata.setter
    def metadata(self, metadata):
        """Sets the metadata of this AdminNamedEntityUpdateRequest.


        :param metadata: The metadata of this AdminNamedEntityUpdateRequest.  # noqa: E501
        :type: AdminNamedEntityMetadata
        """

        self._metadata = metadata

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminNamedEntityUpdateRequest, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminNamedEntityUpdateRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
