import { GetKeysTree } from "../../../wailsjs/go/api/KeyApi";
import { ListData } from "../../../wailsjs/go/api/DataApi"
import { dto } from "../../../wailsjs/go/models";

export const useData = () => {
    let api = {
        getKeyTree: async (req: { id: string, keyword: string }) => {
            const r = await GetKeysTree({id: req.id, keyword: req.keyword, with_suffix: false});
            console.log(r)
            return r;
        },
        listData: async (req: any) => {
            const r = await ListData(req as dto.ListDataRequest);
            console.log(r)
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