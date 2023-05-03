import mitt from 'mitt'
import { ITable, IEndpoint } from '~/lib';

type ApplicationEvents = {
    'sidebar:table:selected': ITable,
    'sidebar:endpoint:selected': IEndpoint,
    'table:deleted': void,
};


export default defineNuxtPlugin(() => {
    const emitter = mitt<ApplicationEvents>()

    return {
        provide: {
            event: emitter.emit, // Will emit an event
            listen: emitter.on // Will register a listener for an event
        }
    }
})