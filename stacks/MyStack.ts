import { StackContext, Api, EventBus, StaticSite } from 'sst/constructs';

export function API({ stack }: StackContext) {
    const bus = new EventBus(stack, 'bus', {
        defaults: {
            retries: 10,
        },
    });

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

    api.addRoutes(stack, {
        'GET /todo': {
            function: {
                runtime: 'go',
                handler: 'backend/cmd/handlers/todo/list/list.go',
            },
        },
    });

    stack.addOutputs({
        ApiEndpoint: api.url,
    });
}
