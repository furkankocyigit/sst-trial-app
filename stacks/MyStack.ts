import { StackContext, Api, EventBus, Table, Function } from 'sst/constructs';

export function API({ stack }: StackContext) {
    const bus = new EventBus(stack, 'bus', {
        defaults: {
            retries: 10,
        },
    });

    //SST-Dynamodb Table
    const todoTable = new Table(stack, 'todoTable', {
        fields: {
            id: 'string',
            title: 'string',
            description: 'string',
        },
        primaryIndex: { partitionKey: 'id' },
    });

    // SST API
    const api = new Api(stack, 'api', {
        defaults: {
            function: {
                bind: [bus],
            },
        },
    });

    // add hello route
    api.addRoutes(stack, {
        'GET /': {
            function: {
                runtime: 'go',
                handler: 'backend/cmd/handlers/hello/hello.go',
            },
        },
    });

    // add todo routes

    api.addRoutes(stack, {
        ['GET /todo']: {
            function: {
                runtime: 'go',
                handler: 'backend/cmd/handlers/todo/list/list.go',
                permissions: [todoTable],
            },
        },
        ['GET /todo/{id}']: {
            function: {
                runtime: 'go',
                handler: 'backend/cmd/handlers/todo/find/find.go',
                permissions: [todoTable],
            },
        },
        ['POST /todo']: {
            function: {
                runtime: 'go',
                handler: 'backend/cmd/handlers/todo/create/create.go',
                permissions: [todoTable],
            },
        },
        ['DELETE /todo/{id}']: {
            function: {
                runtime: 'go',
                handler: 'backend/cmd/handlers/todo/delete/delete.go',
                permissions: [todoTable],
            },
        },
    });

    stack.addOutputs({
        ApiEndpoint: api.url,
        DynamoDBTable: todoTable.tableName,
    });
}
