// model

export interface Connection {
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
}