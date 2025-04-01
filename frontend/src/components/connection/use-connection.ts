import { GetConnection, ListConnections, CreateConnection, UpdateConnection, DeleteConnection, TestConnection } from "../../../wailsjs/go/api/ConnectionApi";
import { dto } from "../../../wailsjs/go/models";

export const useConnection = () => {
    let api = {
        get: async (id: string) => {
            const r = await GetConnection({ id: id });
            console.log(r);
            return r;
        },
        list: async (data: dto.ListConnectionRequest) => {
            const r = await ListConnections(data);
            console.log(r);
            return r;
        },
        create: async (data: dto.CreateConnectionRequest) => {
            const r = await CreateConnection(data);
            console.log(r);
            return r;
        },
        update: async (data: dto.UpdateConnectionRequest) => {
            console.log(data);
            const r = await UpdateConnection(data);
            console.log(r);
            return r;
        },
        delete: async (id: string) => {
            const r = await DeleteConnection({ id: id });
            console.log(r);
            return r;
        },
        test: async (data: dto.TestConnectionRequest) => {
            const r = await TestConnection(data);
            console.log(r);
            return r;
        },
    }

    let fun = {
    }

    return {
        api,
        fun,
    }
}; 