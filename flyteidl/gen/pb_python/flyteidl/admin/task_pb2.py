# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/task.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from flyteidl.core import tasks_pb2 as flyteidl_dot_core_dot_tasks__pb2
from flyteidl.core import compiler_pb2 as flyteidl_dot_core_dot_compiler__pb2
from flyteidl.admin import description_entity_pb2 as flyteidl_dot_admin_dot_description__entity__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19\x66lyteidl/admin/task.proto\x12\x0e\x66lyteidl.admin\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x19\x66lyteidl/core/tasks.proto\x1a\x1c\x66lyteidl/core/compiler.proto\x1a\'flyteidl/admin/description_entity.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"l\n\x11TaskCreateRequest\x12)\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.IdentifierR\x02id\x12,\n\x04spec\x18\x02 \x01(\x0b\x32\x18.flyteidl.admin.TaskSpecR\x04spec\"\x14\n\x12TaskCreateResponse\"\x95\x01\n\x04Task\x12)\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.IdentifierR\x02id\x12\x35\n\x07\x63losure\x18\x02 \x01(\x0b\x32\x1b.flyteidl.admin.TaskClosureR\x07\x63losure\x12+\n\x11short_description\x18\x03 \x01(\tR\x10shortDescription\"L\n\x08TaskList\x12*\n\x05tasks\x18\x01 \x03(\x0b\x32\x14.flyteidl.admin.TaskR\x05tasks\x12\x14\n\x05token\x18\x02 \x01(\tR\x05token\"\x88\x01\n\x08TaskSpec\x12\x37\n\x08template\x18\x01 \x01(\x0b\x32\x1b.flyteidl.core.TaskTemplateR\x08template\x12\x43\n\x0b\x64\x65scription\x18\x02 \x01(\x0b\x32!.flyteidl.admin.DescriptionEntityR\x0b\x64\x65scription\"\x8a\x01\n\x0bTaskClosure\x12@\n\rcompiled_task\x18\x01 \x01(\x0b\x32\x1b.flyteidl.core.CompiledTaskR\x0c\x63ompiledTask\x12\x39\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAtB\xaf\x01\n\x12\x63om.flyteidl.adminB\tTaskProtoP\x01Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin\xa2\x02\x03\x46\x41X\xaa\x02\x0e\x46lyteidl.Admin\xca\x02\x0e\x46lyteidl\\Admin\xe2\x02\x1a\x46lyteidl\\Admin\\GPBMetadata\xea\x02\x0f\x46lyteidl::Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.admin.task_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.flyteidl.adminB\tTaskProtoP\001Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin\242\002\003FAX\252\002\016Flyteidl.Admin\312\002\016Flyteidl\\Admin\342\002\032Flyteidl\\Admin\\GPBMetadata\352\002\017Flyteidl::Admin'
  _globals['_TASKCREATEREQUEST']._serialized_start=208
  _globals['_TASKCREATEREQUEST']._serialized_end=316
  _globals['_TASKCREATERESPONSE']._serialized_start=318
  _globals['_TASKCREATERESPONSE']._serialized_end=338
  _globals['_TASK']._serialized_start=341
  _globals['_TASK']._serialized_end=490
  _globals['_TASKLIST']._serialized_start=492
  _globals['_TASKLIST']._serialized_end=568
  _globals['_TASKSPEC']._serialized_start=571
  _globals['_TASKSPEC']._serialized_end=707
  _globals['_TASKCLOSURE']._serialized_start=710
  _globals['_TASKCLOSURE']._serialized_end=848
# @@protoc_insertion_point(module_scope)
