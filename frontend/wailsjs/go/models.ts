export namespace dto {
	
	export class BaseResponse {
	    success: boolean;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new BaseResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}
	export class CreateConnectionRequest {
	    id: string;
	    name: string;
	    host: string;
	    port: number;
	    enable_auth: boolean;
	    username: string;
	    password: string;
	    enable_tls: boolean;
	    ca: string;
	    cert: string;
	    key: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateConnectionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.enable_auth = source["enable_auth"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.enable_tls = source["enable_tls"];
	        this.ca = source["ca"];
	        this.cert = source["cert"];
	        this.key = source["key"];
	    }
	}
	export class DeleteConnectionRequest {
	    id: string;
	
	    static createFrom(source: any = {}) {
	        return new DeleteConnectionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	export class DeleteDataRequest {
	    id: string;
	    key: string;
	
	    static createFrom(source: any = {}) {
	        return new DeleteDataRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.key = source["key"];
	    }
	}
	export class GetConnectionRequest {
	    id: string;
	
	    static createFrom(source: any = {}) {
	        return new GetConnectionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	export class GetKeysTreeRequest {
	    id: string;
	    keyword: string;
	    with_suffix: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GetKeysTreeRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.keyword = source["keyword"];
	        this.with_suffix = source["with_suffix"];
	    }
	}
	export class ListConnectionRequest {
	    page: number;
	    size: number;
	    total: number;
	    keyword: string;
	
	    static createFrom(source: any = {}) {
	        return new ListConnectionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.size = source["size"];
	        this.total = source["total"];
	        this.keyword = source["keyword"];
	    }
	}
	export class ListDataParams {
	    id: string;
	    prefix: string;
	
	    static createFrom(source: any = {}) {
	        return new ListDataParams(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.prefix = source["prefix"];
	    }
	}
	export class ListDataRequest {
	    page: number;
	    size: number;
	    total: number;
	    keyword: string;
	    params: ListDataParams;
	
	    static createFrom(source: any = {}) {
	        return new ListDataRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.page = source["page"];
	        this.size = source["size"];
	        this.total = source["total"];
	        this.keyword = source["keyword"];
	        this.params = this.convertValues(source["params"], ListDataParams);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PutDataRequest {
	    id: string;
	    key: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new PutDataRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class TestConnectionRequest {
	    id: string;
	    name: string;
	    host: string;
	    port: number;
	    enable_auth: boolean;
	    username: string;
	    password: string;
	    enable_tls: boolean;
	    ca: string;
	    cert: string;
	    key: string;
	
	    static createFrom(source: any = {}) {
	        return new TestConnectionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.enable_auth = source["enable_auth"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.enable_tls = source["enable_tls"];
	        this.ca = source["ca"];
	        this.cert = source["cert"];
	        this.key = source["key"];
	    }
	}
	export class UpdateConnectionRequest {
	    id: string;
	    name: string;
	    host: string;
	    port: number;
	    enable_auth: boolean;
	    username: string;
	    password: string;
	    enable_tls: boolean;
	    ca: string;
	    cert: string;
	    key: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateConnectionRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.enable_auth = source["enable_auth"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.enable_tls = source["enable_tls"];
	        this.ca = source["ca"];
	        this.cert = source["cert"];
	        this.key = source["key"];
	    }
	}

}

