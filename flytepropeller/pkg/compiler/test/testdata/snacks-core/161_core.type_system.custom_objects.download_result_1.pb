
�
3"/core.type_system.custom_objects.download_resultpython-task0.32.6python* "�
�
�
res�
�	2�
$
$ref#/definitions/ResultSchema
4
$schema)'http://json-schema.org/draft-07/schema#
�
definitions�*�
�
FlyteschemaSchema*}

typeobject
M

properties?*=
;
remote_path,**

typestring

titleremote_path

additionalProperties  
�
FlytefileSchemaq*o

typeobject
?

properties1*/
-
path%*#

typestring

titlepath

additionalProperties  
�
FlytedirectorySchemaq*o
?

properties1*/
-
path%*#

typestring

titlepath

additionalProperties  

typeobject
�
ResultSchema�*�

additionalProperties  
�

properties�*�
W
fileO*M

typeobject
'
$ref#/definitions/FlytefileSchema


field_many  
[
schemaQ*O


field_many  

typeobject
)
$ref!#/definitions/FlyteschemaSchema
a
	directoryT*R
,
$ref$"#/definitions/FlytedirectorySchema

typeobject


field_many  

typeobjectres 2�
Lghcr.io/flyteorg/flytecookbook:core-8b8e1a849c9adfca88049a074b10dad278f70077pyflyte-execute--inputs
{{.input}}--output-prefix{{.outputPrefix}}--raw-output-data-prefix{{.rawOutputDataPrefix}}--checkpoint-path{{.checkpointOutputPrefix}}--prev-checkpoint{{.prevCheckpointPrefix}}
--resolver9flytekit.core.python_auto_container.default_task_resolver--task-modulecore.type_system.custom_objects	task-namedownload_result" 