import { ApiHandler } from 'sst/node/api';

export const handler = ApiHandler(async (_evt) => {
    return {
        statusCode: 200,
        body: `Created by SST console ${new Date().toISOString()}`,
    };
});
