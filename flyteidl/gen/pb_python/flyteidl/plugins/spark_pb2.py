# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/plugins/spark.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from flyteidl.core import tasks_pb2 as flyteidl_dot_core_dot_tasks__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x66lyteidl/plugins/spark.proto\x12\x10\x66lyteidl.plugins\x1a\x1cgoogle/protobuf/struct.proto\x1a\x19\x66lyteidl/core/tasks.proto\"B\n\x10SparkApplication\".\n\x04Type\x12\n\n\x06PYTHON\x10\x00\x12\x08\n\x04JAVA\x10\x01\x12\t\n\x05SCALA\x10\x02\x12\x05\n\x01R\x10\x03\"\xec\x05\n\x08SparkJob\x12Q\n\x0f\x61pplicationType\x18\x01 \x01(\x0e\x32\'.flyteidl.plugins.SparkApplication.TypeR\x0f\x61pplicationType\x12\x30\n\x13mainApplicationFile\x18\x02 \x01(\tR\x13mainApplicationFile\x12\x1c\n\tmainClass\x18\x03 \x01(\tR\tmainClass\x12G\n\tsparkConf\x18\x04 \x03(\x0b\x32).flyteidl.plugins.SparkJob.SparkConfEntryR\tsparkConf\x12J\n\nhadoopConf\x18\x05 \x03(\x0b\x32*.flyteidl.plugins.SparkJob.HadoopConfEntryR\nhadoopConf\x12\"\n\x0c\x65xecutorPath\x18\x06 \x01(\tR\x0c\x65xecutorPath\x12?\n\x0e\x64\x61tabricksConf\x18\x07 \x01(\x0b\x32\x17.google.protobuf.StructR\x0e\x64\x61tabricksConf\x12(\n\x0f\x64\x61tabricksToken\x18\x08 \x01(\tR\x0f\x64\x61tabricksToken\x12.\n\x12\x64\x61tabricksInstance\x18\t \x01(\tR\x12\x64\x61tabricksInstance\x12\x33\n\tdriverPod\x18\n \x01(\x0b\x32\x15.flyteidl.core.K8sPodR\tdriverPod\x12\x37\n\x0b\x65xecutorPod\x18\x0b \x01(\x0b\x32\x15.flyteidl.core.K8sPodR\x0b\x65xecutorPod\x1a<\n\x0eSparkConfEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x1a=\n\x0fHadoopConfEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x42\xc2\x01\n\x14\x63om.flyteidl.pluginsB\nSparkProtoP\x01Z=github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/plugins\xa2\x02\x03\x46PX\xaa\x02\x10\x46lyteidl.Plugins\xca\x02\x10\x46lyteidl\\Plugins\xe2\x02\x1c\x46lyteidl\\Plugins\\GPBMetadata\xea\x02\x11\x46lyteidl::Pluginsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.plugins.spark_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\024com.flyteidl.pluginsB\nSparkProtoP\001Z=github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/plugins\242\002\003FPX\252\002\020Flyteidl.Plugins\312\002\020Flyteidl\\Plugins\342\002\034Flyteidl\\Plugins\\GPBMetadata\352\002\021Flyteidl::Plugins'
  _SPARKJOB_SPARKCONFENTRY._options = None
  _SPARKJOB_SPARKCONFENTRY._serialized_options = b'8\001'
  _SPARKJOB_HADOOPCONFENTRY._options = None
  _SPARKJOB_HADOOPCONFENTRY._serialized_options = b'8\001'
  _globals['_SPARKAPPLICATION']._serialized_start=107
  _globals['_SPARKAPPLICATION']._serialized_end=173
  _globals['_SPARKAPPLICATION_TYPE']._serialized_start=127
  _globals['_SPARKAPPLICATION_TYPE']._serialized_end=173
  _globals['_SPARKJOB']._serialized_start=176
  _globals['_SPARKJOB']._serialized_end=924
  _globals['_SPARKJOB_SPARKCONFENTRY']._serialized_start=801
  _globals['_SPARKJOB_SPARKCONFENTRY']._serialized_end=861
  _globals['_SPARKJOB_HADOOPCONFENTRY']._serialized_start=863
  _globals['_SPARKJOB_HADOOPCONFENTRY']._serialized_end=924
# @@protoc_insertion_point(module_scope)
