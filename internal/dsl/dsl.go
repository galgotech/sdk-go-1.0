package dsl

const DSLSpec = `
$id: https://serverlessworkflow.io/schemas/1.0.0-alpha1/workflow.yaml
$schema: https://json-schema.org/draft/2020-12/schema
description: Serverless Workflow DSL - Workflow Schema
type: object
properties:
  document:
    type: object
    properties:
      dsl:
        type: string
        pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
        description: The version of the DSL used by the workflow.
      namespace:
        type: string
        pattern: ^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$
        description: The workflow's namespace.
      name:
        type: string
        pattern: ^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$
        description: The workflow's name.
      version:
        type: string
        pattern: ^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
        description: The workflow's semantic version.
      title:
        type: string
        description: The workflow's title.
      summary:
        type: string
        description: The workflow's Markdown summary.
      tags:
        type: object
        description: A key/value mapping of the workflow's tags, if any.
        additionalProperties: true
    required: [ dsl, namespace, name, version ]
    description: Documents the workflow
  input:
    $ref: '#/$defs/input'
    description: Configures the workflow's input.
  use:
    type: object
    properties:
      authentications:
        type: object
        additionalProperties:
          $ref: '#/$defs/authenticationPolicy'
        description: The workflow's reusable authentication policies.
      errors:
        type: object
        additionalProperties:
          $ref: '#/$defs/error'
        description: The workflow's reusable errors.
      extensions:
        type: array
        items:
          type: object
          title: ExtensionItem
          minProperties: 1
          maxProperties: 1
          additionalProperties:
            $ref: '#/$defs/extension'
        description: The workflow's extensions.
      functions:
        type: object
        additionalProperties:
          $ref: '#/$defs/task'
        description: The workflow's reusable functions.
      retries:
        type: object
        additionalProperties:
          $ref: '#/$defs/retryPolicy'
        description: The workflow's reusable retry policies.
      secrets:
        type: array
        items:
          type: string
        description: The workflow's secrets.
    description: Defines the workflow's reusable components.
  do:
    description: Defines the task(s) the workflow must perform
    $ref: '#/$defs/taskList'
  timeout:
    $ref: '#/$defs/timeout'
    description: The workflow's timeout configuration, if any.
  output:
    $ref: '#/$defs/output'
    description: Configures the workflow's output.
  schedule:
    type: object
    properties:
      every:
        $ref: '#/$defs/duration'
        description: Specifies the duration of the interval at which the workflow should be executed.
      cron:
        type: string
        description: Specifies the schedule using a cron expression, e.g., '0 0 * * *' for daily at midnight."
      after:
        $ref: '#/$defs/duration'
        description: Specifies a delay duration that the workflow must wait before starting again after it completes.
      on:
        $ref: '#/$defs/eventConsumptionStrategy'
        description: Specifies the events that trigger the workflow execution.
    description: Schedules the workflow
$defs:
  taskList:
    type: array
    items:
      type: object
      title: TaskItem
      minProperties: 1
      maxProperties: 1
      additionalProperties:
        $ref: '#/$defs/task'
  taskBase:
    type: object
    properties:
      if:
        type: string
        description: A runtime expression, if any, used to determine whether or not the task should be run.
      input:
        $ref: '#/$defs/input'
        description: Configure the task's input.
      output:
        $ref: '#/$defs/output'
        description: Configure the task's output.
      export:
        $ref: '#/$defs/export'
        description: Export task output to context.
      timeout:
        $ref: '#/$defs/timeout'
        description: The task's timeout configuration, if any.
      then:
        $ref: '#/$defs/flowDirective'
        description: The flow directive to be performed upon completion of the task.
  task:
    unevaluatedProperties: false
    oneOf:
      - $ref: '#/$defs/callTask'
      - $ref: '#/$defs/doTask'
      - $ref: '#/$defs/forkTask'
      - $ref: '#/$defs/emitTask'
      - $ref: '#/$defs/forTask'
      - $ref: '#/$defs/listenTask'
      - $ref: '#/$defs/raiseTask'
      - $ref: '#/$defs/runTask'
      - $ref: '#/$defs/setTask'
      - $ref: '#/$defs/switchTask'
      - $ref: '#/$defs/tryTask'
      - $ref: '#/$defs/waitTask'
  callTask:
    oneOf:
      - title: CallAsyncAPI
        $ref: '#/$defs/taskBase'
        type: object
        required: [ call, with ]
        unevaluatedProperties: false
        properties:
          call:
            type: string
            const: asyncapi
          with:
            title: WithAsyncAPI
            type: object
            properties:
              document:
                $ref: '#/$defs/externalResource'
                description: The document that defines the AsyncAPI operation to call.
              operationRef:
                type: string
                description: A reference to the AsyncAPI operation to call.
              server:
                type: string
                description: A a reference to the server to call the specified AsyncAPI operation on. If not set, default to the first server matching the operation's channel.
              message:
                type: string
                description: The name of the message to use. If not set, defaults to the first message defined by the operation.
              binding:
                type: string
                description: The name of the binding to use. If not set, defaults to the first binding defined by the operation.
              payload:
                type: object
                description: The payload to call the AsyncAPI operation with, if any.
              authentication:
                $ref: '#/$defs/referenceableAuthenticationPolicy'
                description: The authentication policy, if any, to use when calling the AsyncAPI operation.
            required: [ document, operationRef ]
            additionalProperties: false
            description: Defines the AsyncAPI call to perform.
      - title: CallGRPC
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call, with ]
        properties:
          call:
            type: string
            const: grpc
          with:
            title: WithGRPC
            type: object
            properties:
              proto:
                $ref: '#/$defs/externalResource'
                description: The proto resource that describes the GRPC service to call.
              service:
                type: object
                properties:
                  name:
                    type: string
                    description: The name of the GRPC service to call.
                  host:
                    type: string
                    pattern: ^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$
                    description: The hostname of the GRPC service to call.
                  port:
                    type: integer
                    min: 0
                    max: 65535
                    description: The port number of the GRPC service to call.
                  authentication:
                    $ref: '#/$defs/referenceableAuthenticationPolicy'
                    description: The endpoint's authentication policy, if any.
                required: [ name, host ]
              method:
                type: string
                description: The name of the method to call on the defined GRPC service.
              arguments:
                type: object
                additionalProperties: true
                description: The arguments, if any, to call the method with.
            required: [ proto, service, method ]
            additionalProperties: false
            description: Defines the GRPC call to perform.
      - title: CallHTTP
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call, with ]
        properties:
          call:
            type: string
            const: http
          with:
            title: WithHTTP
            type: object
            properties:
              method:
                type: string
                description: The HTTP method of the HTTP request to perform.
              endpoint:
                description: The HTTP endpoint to send the request to.
                oneOf:
                  - $ref: '#/$defs/endpoint'
                  - type: string
                    format: uri-template
              headers:
                type: object
                description: A name/value mapping of the headers, if any, of the HTTP request to perform.
              body:
                description: The body, if any, of the HTTP request to perform.
              output:
                type: string
                enum: [ raw, content, response ]
                description: The http call output format. Defaults to 'content'.
            required: [ method, endpoint ]
            additionalProperties: false
            description: Defines the HTTP call to perform.
      - title: CallOpenAPI
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call, with ]
        properties:
          call:
            type: string
            const: openapi
          with:
            title: WithOpenAPI
            type: object
            properties:
              document:
                $ref: '#/$defs/externalResource'
                description: The document that defines the OpenAPI operation to call.
              operationId:
                type: string
                description: The id of the OpenAPI operation to call.
              parameters:
                type: object
                additionalProperties: true
                description: A name/value mapping of the parameters of the OpenAPI operation to call.
              authentication:
                $ref: '#/$defs/referenceableAuthenticationPolicy'
                description: The authentication policy, if any, to use when calling the OpenAPI operation.
              output:
                type: string
                enum: [ raw, content, response ]
                description: The http call output format. Defaults to 'content'.
            required: [ document, operationId ]
            additionalProperties: false
            description: Defines the OpenAPI call to perform.
      - title: CallFunction
        $ref: '#/$defs/taskBase'
        type: object
        unevaluatedProperties: false
        required: [ call ]
        properties:
          call:
            type: string
            not:
              enum: ["asyncapi", "grpc", "http", "openapi"]
            description: The name of the function to call.
          with:
            type: object
            additionalProperties: true
            description: A name/value mapping of the parameters, if any, to call the function with.
  forkTask:
    description: Allows workflows to execute multiple tasks concurrently and optionally race them against each other, with a single possible winner, which sets the task's output.
    $ref: '#/$defs/taskBase'
    type: object
    unevaluatedProperties: false
    required: [ fork ]
    properties:
      fork:
        type: object
        required: [ branches ]
        properties:
          branches:
            $ref: '#/$defs/taskList'
          compete:
            description: Indicates whether or not the concurrent tasks are racing against each other, with a single possible winner, which sets the composite task's output.
            type: boolean
            default: false  
  doTask:
    description: Allows to execute a list of tasks in sequence
    $ref: '#/$defs/taskBase'
    type: object
    unevaluatedProperties: false
    required: [ do ]
    properties:
      do:
        $ref: '#/$defs/taskList'
  emitTask:
    description: Allows workflows to publish events to event brokers or messaging systems, facilitating communication and coordination between different components and services.
    $ref: '#/$defs/taskBase'
    type: object
    required: [ emit ]
    unevaluatedProperties: false
    properties:
      emit:
        type: object
        properties:
          event:
            type: object
            properties:
              id:
                type: string
                description: The event's unique identifier
              source:
                type: string
                format: uri
                description: Identifies the context in which an event happened
              type:
                type: string
                description: This attribute contains a value describing the type of event related to the originating occurrence.
              time:
                type: string
                format: date-time
              subject:
                type: string
              datacontenttype:
                type: string
                description: Content type of data value. This attribute enables data to carry any type of content, whereby format and encoding might differ from that of the chosen event format.
              dataschema:
                type: string
                format: uri
            required: [ source, type ]
            additionalProperties: true
        required: [ event ]
  forTask:
    description: Allows workflows to iterate over a collection of items, executing a defined set of subtasks for each item in the collection. This task type is instrumental in handling scenarios such as batch processing, data transformation, and repetitive operations across datasets.
    $ref: '#/$defs/taskBase'
    type: object
    required: [ for, do ]
    unevaluatedProperties: false
    properties:
      for:
        type: object
        properties:
          each:
            type: string
            description: The name of the variable used to store the current item being enumerated.
            default: item
          in:
            type: string
            description: A runtime expression used to get the collection to enumerate.
          at:
            type: string
            description: The name of the variable used to store the index of the current item being enumerated.
            default: index
        required: [ in ]
      while:
        type: string
        description: A runtime expression that represents the condition, if any, that must be met for the iteration to continue.
      do:
        $ref: '#/$defs/taskList'
  listenTask:
    description: Provides a mechanism for workflows to await and react to external events, enabling event-driven behavior within workflow systems.
    $ref: '#/$defs/taskBase'
    type: object
    required: [ listen ]
    unevaluatedProperties: false
    properties:
      listen:
        type: object
        properties:
          to:
            $ref: '#/$defs/eventConsumptionStrategy'
            description: Defines the event(s) to listen to.
        required: [ to ]
  raiseTask:
    description: Intentionally triggers and propagates errors.
    $ref: '#/$defs/taskBase'
    type: object
    required: [ raise ]
    unevaluatedProperties: false
    properties:
      raise:
        type: object
        properties:
          error:
            $ref: '#/$defs/error'
            description: Defines the error to raise.
        required: [ error ]
  runTask:
    description: Provides the capability to execute external containers, shell commands, scripts, or workflows.
    $ref: '#/$defs/taskBase'
    type: object
    required: [ run ]
    unevaluatedProperties: false
    properties:
      run:
        type: object
        oneOf:
          - title: RunContainer
            properties:
              container:
                type: object
                properties:
                  image:
                    type: string
                    description: The name of the container image to run.
                  command:
                    type: string
                    description: The command, if any, to execute on the container
                  ports:
                    type: object
                    description: The container's port mappings, if any.
                  volumes:
                    type: object
                    description: The container's volume mappings, if any.
                  environment:
                    title: ContainerEnvironment
                    type: object
                    description: A key/value mapping of the environment variables, if any, to use when running the configured process.
                required: [ image ]
            required: [ container ]
            description: Enables the execution of external processes encapsulated within a containerized environment.
          - title: RunScript
            properties:
              script:
                type: object
                properties:
                  language:
                    type: string
                    description: The language of the script to run.
                  environment:
                    title: ScriptEnvironment
                    type: object
                    additionalProperties: true
                    description: A key/value mapping of the environment variables, if any, to use when running the configured process.
                oneOf:
                  - title: ScriptInline
                    properties:
                      code:
                        type: string
                    required: [ code ]
                    description: The script's code.
                  - title: ScriptExternal
                    properties:
                      source:
                        $ref: '#/$defs/externalResource'
                    description: The script's resource.
                    required: [ source ]
                required: [ language ]
            required: [ script ]
            description: Enables the execution of custom scripts or code within a workflow, empowering workflows to perform specialized logic, data processing, or integration tasks by executing user-defined scripts written in various programming languages.
          - title: RunShell
            properties:
              shell:
                type: object
                properties:
                  command:
                    type: string
                    description: The shell command to run.
                  arguments:
                    title: ShellArguments
                    type: object
                    additionalProperties: true
                    description: A list of the arguments of the shell command to run.
                  environment:
                    title: ShellEnvironment
                    type: object
                    additionalProperties: true
                    description: A key/value mapping of the environment variables, if any, to use when running the configured process.
                required: [ command ]
            required: [ shell ]
            description: Enables the execution of shell commands within a workflow, enabling workflows to interact with the underlying operating system and perform system-level operations, such as file manipulation, environment configuration, or system administration tasks.
          - title: RunWokflow
            properties:
              workflow:
                title: RunWorkflowDescriptor
                type: object
                properties:
                  namespace:
                    type: string
                    description: The namespace the workflow to run belongs to.
                  name:
                    type: string
                    description: The name of the workflow to run.
                  version:
                    type: string
                    default: latest
                    description: The version of the workflow to run. Defaults to latest
                  input:
                    title: WorkflowInput
                    type: object
                    additionalProperties: true
                    description: The data, if any, to pass as input to the workflow to execute. The value should be validated against the target workflow's input schema, if specified.
                required: [ namespace, name, version ]
            required: [ workflow ]
            description: Enables the invocation and execution of nested workflows within a parent workflow, facilitating modularization, reusability, and abstraction of complex logic or business processes by encapsulating them into standalone workflow units.
  setTask:
    description: A task used to set data
    $ref: '#/$defs/taskBase'
    type: object
    required: [ set ]
    unevaluatedProperties: false
    properties:
      set:
        type: object
        minProperties: 1
        additionalProperties: true
        description: The data to set
  switchTask:
    description: Enables conditional branching within workflows, allowing them to dynamically select different paths based on specified conditions or criteria
    $ref: '#/$defs/taskBase'
    type: object
    required: [ switch ]
    unevaluatedProperties: false
    properties:
      switch:
        type: array
        minItems: 1
        items:
          type: object
          minProperties: 1
          maxProperties: 1
          title: SwitchItem
          additionalProperties:
            type: object
            title: SwitchCase
            properties:
              name:
                type: string
                description: The case's name.
              when:
                type: string
                description: A runtime expression used to determine whether or not the case matches.
              then:
                $ref: '#/$defs/flowDirective'
                description: The flow directive to execute when the case matches.
  tryTask:
    description: Serves as a mechanism within workflows to handle errors gracefully, potentially retrying failed tasks before proceeding with alternate ones.
    $ref: '#/$defs/taskBase'
    type: object
    required: [ try, catch ]
    unevaluatedProperties: false
    properties:
      try:
        description: The task(s) to perform.
        $ref: '#/$defs/taskList'
      catch:
        type: object
        properties:
          errors:
            title: CatchErrors
            type: object
          as:
            type: string
            description: The name of the runtime expression variable to save the error as. Defaults to 'error'.
          when:
            type: string
            description: A runtime expression used to determine whether or not to catch the filtered error
          exceptWhen:
            type: string
            description: A runtime expression used to determine whether or not to catch the filtered error
          retry:
            $ref: '#/$defs/retryPolicy'
            description: The retry policy to use, if any, when catching errors.
          do:
            description: The definition of the task(s) to run when catching an error.
            $ref: '#/$defs/taskList'
  waitTask:
    description: Allows workflows to pause or delay their execution for a specified period of time.
    $ref: '#/$defs/taskBase'
    type: object
    required: [ wait ]
    unevaluatedProperties: false
    properties:
      wait:
        description: The amount of time to wait.
        $ref: '#/$defs/duration'
  flowDirective:
    additionalProperties: false
    anyOf:
      - type: string
        enum: [ continue, exit, end ]
        default: continue
      - type: string
  referenceableAuthenticationPolicy:
    type: object
    oneOf:
      - title: AuthenticationPolicyReference
        properties:
          use:
            type: string
            minLength: 1
            description: The name of the authentication policy to use
        required: [use]
      - $ref: '#/$defs/authenticationPolicy'
  secretBasedAuthenticationPolicy:
    type: object
    properties:
      use:
        type: string
        minLength: 1
        description: The name of the authentication policy to use
    required: [use]
  authenticationPolicy:
    type: object
    oneOf:
    - title: BasicAuthenticationPolicy 
      properties:
        basic:
          type: object
          oneOf:
            - properties:
                username:
                  type: string
                  description: The username to use.
                password:
                  type: string
                  description: The password to use.
              required: [ username, password ]
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
      required: [ basic ]
      description: Use basic authentication.
    - title: BearerAuthenticationPolicy
      properties:
        bearer:
          type: object
          oneOf:
            - properties:
                token:
                  type: string
                  description: The bearer token to use.
              required: [ token ]
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
      required: [ bearer ]
      description: Use bearer authentication.
    - title: OAuth2AuthenticationPolicy
      properties:
        oauth2:
          type: object
          oneOf:
            - properties:
                authority:
                  type: string
                  format: uri
                  description: The URI that references the OAuth2 authority to use.
                grant:
                  type: string
                  description: The grant type to use.
                client:
                  type: object
                  properties:
                    id:
                      type: string
                      description: The client id to use.
                    secret:
                      type: string
                      description: The client secret to use, if any.
                  required: [ id ]
                scopes:
                  type: array
                  items:
                    type: string
                  description: The scopes, if any, to request the token for.
                audiences:
                  type: array
                  items:
                    type: string
                  description: The audiences, if any, to request the token for.
                username:
                  type: string
                  description: The username to use. Used only if the grant type is Password.
                password:
                  type: string
                  description: The password to use. Used only if the grant type is Password.
                subject:
                  $ref: '#/$defs/oauth2Token'
                  description: The security token that represents the identity of the party on behalf of whom the request is being made.
                actor:
                  $ref: '#/$defs/oauth2Token'
                  description: The security token that represents the identity of the acting party.
              required: [ authority, grant, client ]
            - $ref: '#/$defs/secretBasedAuthenticationPolicy'
      required: [ oauth2 ]
      description: Use OAUTH2 authentication.
    description: Defines an authentication policy.
  oauth2Token:
    type: object
    properties:
      token:
        type: string
        description: The security token to use to use.
      type:
        type: string
        description: The type of the security token to use to use.
    required: [ token, type ]
  duration:
    type: object
    minProperties: 1
    properties:
      days:
        type: integer
        description: Number of days, if any.
      hours:
        type: integer
        description: Number of days, if any.
      minutes:
        type: integer
        description: Number of minutes, if any.
      seconds:
        type: integer
        description: Number of seconds, if any.
      milliseconds:
        type: integer
        description: Number of milliseconds, if any.
    description: The definition of a duration.
  error:
    type: object
    properties:
      type:
        type: string
        format: uri
        description: A URI reference that identifies the error type.
      status:
        type: integer
        description: The status code generated by the origin for this occurrence of the error.
      instance:
        type: string
        format: json-pointer
        description: A JSON Pointer used to reference the component the error originates from.
      title:
        type: string
        description: A short, human-readable summary of the error.
      detail:
        type: string
        description: A human-readable explanation specific to this occurrence of the error.
    required: [ type, status, instance ]
  endpoint:
    type: object
    properties:
      uri:
        type: string
        format: uri-template
        description: The endpoint's URI.
      authentication:
        $ref: '#/$defs/referenceableAuthenticationPolicy'
        description: The authentication policy to use.
    required: [ uri ]
  eventConsumptionStrategy:
    type: object
    oneOf:
      - title: AllEventConsumptionStrategy
        properties:
          all:
            type: array
            items:
              $ref: '#/$defs/eventFilter'
            description: A list containing all the events that must be consumed.
        required: [ all ]
      - title: AnyEventConsumptionStrategy
        properties:
          any:
            type: array
            items:
              $ref: '#/$defs/eventFilter'
            description: A list containing any of the events to consume.
        required: [ any ]
      - title: OneEventConsumptionStrategy
        properties:
          one:
            $ref: '#/$defs/eventFilter'
            description: The single event to consume.
        required: [ one ]
  eventFilter:
    type: object
    properties:
      with:
        title: WithEvent
        type: object
        minProperties: 1
        properties:
          id:
            type: string
            description: The event's unique identifier
          source:
            type: string
            description: Identifies the context in which an event happened
          type:
            type: string
            description: This attribute contains a value describing the type of event related to the originating occurrence.
          time:
            type: string
          subject:
            type: string
          datacontenttype:
            type: string
            description: Content type of data value. This attribute enables data to carry any type of content, whereby format and encoding might differ from that of the chosen event format.
          dataschema:
            type: string
        additionalProperties: true
        description: An event filter is a mechanism used to selectively process or handle events based on predefined criteria, such as event type, source, or specific attributes.
      correlate:
        type: object
        additionalProperties:
          type: object
          properties:
            from:
              type: string
              description: A runtime expression used to extract the correlation value from the filtered event.
            expect:
              type: string
              description: A constant or a runtime expression, if any, used to determine whether or not the extracted correlation value matches expectations. If not set, the first extracted value will be used as the correlation's expectation.
          required: [ from ]
        description: A correlation is a link between events and data, established by mapping event attributes to specific data attributes, allowing for coordinated processing or handling based on event characteristics.
    required: [ with ]
    description: An event filter is a mechanism used to selectively process or handle events based on predefined criteria, such as event type, source, or specific attributes.
  extension:
    type: object
    properties:
      extend:
        type: string
        enum: [ call, composite, emit, for, listen, raise, run, set, switch, try, wait, all ]
        description: The type of task to extend.
      when:
        type: string
        description: A runtime expression, if any, used to determine whether or not the extension should apply in the specified context.
      before:
        description: The task(s) to execute before the extended task, if any.
        $ref: '#/$defs/taskList'
      after:
        description: The task(s) to execute after the extended task, if any.
        $ref: '#/$defs/taskList'
    required: [ extend ]
    description: The definition of a an extension.
  externalResource:
    oneOf:
      - type: string
        format: uri
      - title: ExternalResourceURI
        type: object
        properties:
          uri:
            type: string
            format: uri
            description: The endpoint's URI.
          authentication:
            $ref: '#/$defs/referenceableAuthenticationPolicy'
            description: The authentication policy to use.
          name:
            type: string
            description: The external resource's name, if any.
        required: [ uri ]
  input:
    type: object
    properties:
      schema:
        $ref: '#/$defs/schema'
        description: The schema used to describe and validate the input of the workflow or task.
      from:
        oneOf:
          - type: string
          - type: object
        description: A runtime expression, if any, used to mutate and/or filter the input of the workflow or task.
    description: Configures the input of a workflow or task.
  output:
    type: object
    properties:
      schema:
        $ref: '#/$defs/schema'
        description: The schema used to describe and validate the output of the workflow or task.
      as:
        oneOf:
          - type: string
          - type: object
        description: A runtime expression, if any, used to mutate and/or filter the output of the workflow or task.
    description: Configures the output of a workflow or task.
  export:
    type: object
    properties:
      schema:
        $ref: '#/$defs/schema'
        description: The schema used to describe and validate the workflow context.
      as:
        oneOf:
          - type: string
          - type: object
        description: A runtime expression, if any, used to export the output data to the context.
    description: Set the content of the context. 
  retryPolicy:
    type: object
    properties:
      when:
        type: string
        description: A runtime expression, if any, used to determine whether or not to retry running the task, in a given context.
      exceptWhen:
        type: string
        description: A runtime expression used to determine whether or not to retry running the task, in a given context.
      delay:
        $ref: '#/$defs/duration'
        description: The duration to wait between retry attempts.
      backoff:
        type: object
        oneOf:
        - title: ConstantBackoff
          properties:
            constant:
              type: object
              description: The definition of the constant backoff to use, if any.
          required: [ constant ]
        - title: ExponentialBackOff
          properties:
            exponential:
              type: object
              description: The definition of the exponential backoff to use, if any.
          required: [ exponential ]
        - title: LinearBackoff
          properties:
            linear:
              type: object
              description: The definition of the linear backoff to use, if any.
          required: [ linear ]
        description: The retry duration backoff.
      limit:
        type: object
        properties:
          attempt:
            type: object
            properties:
              count:
                type: integer
                description: The maximum amount of retry attempts, if any.
              duration:
                $ref: '#/$defs/duration'
                description: The maximum duration for each retry attempt.
          duration:
            $ref: '#/$defs/duration'
            description: The duration limit, if any, for all retry attempts.
        description: The retry limit, if any
      jitter:
        type: object
        properties:
          from:
            $ref: '#/$defs/duration'
            description: The minimum duration of the jitter range
          to:
            $ref: '#/$defs/duration'
            description: The maximum duration of the jitter range
        required: [ from, to ]
        description: The parameters, if any, that control the randomness or variability of the delay between retry attempts.
    description: Defines a retry policy.
  schema:
    type: object
    properties:
      format:
        type: string
        default: json
        description: The schema's format. Defaults to 'json'. The (optional) version of the format can be set using ` + "{format}:{version}" + `.
    oneOf:
      - title: SchemaInline 
        properties:
          document:
            description: The schema's inline definition.
        required: [ document ]
      - title: SchemaExternal
        properties:
          resource:
            $ref: '#/$defs/externalResource'
            description: The schema's external resource.
        required: [ resource ]
    description: Represents the definition of a schema.
  timeout:
    type: object
    properties:
      after:
        $ref: '#/$defs/duration'
        description: The duration after which to timeout.
    required: [ after ]
    description: The definition of a timeout.
required: [ document, do ]
`
