import type { RepoRecord } from "./ repo-types";
import type { CollectionsListResponse, LoginResponse } from "./pb-types";
import { goto } from '$app/navigation';

var _apiClient: IAPIClient;

export function GetApiClient(): IAPIClient {
    if (!_apiClient) {
        _apiClient = new APIClient();
    }

    return _apiClient;
}

export interface IAPIClient {
    TryUseValidToken(): Promise<boolean>;
    Login(username: string, password: string, superuser: boolean): Promise<boolean>;
    ListAllRepos(): Promise<RepoRecord[]>;
}

class APIClient implements IAPIClient {
    constructor() {
        const existingToken = localStorage.getItem("JWT");
        if (existingToken) {
            this.jwtToken = existingToken;
        }
    }

    private baseUrl: string = "http://127.0.0.1:8090"
    private jwtToken: string = "";
    private isSuperUser: boolean = false;

    public async TryUseValidToken(): Promise<boolean> {
        if (!this.jwtToken) {
            return false;
        }

        const res = await fetch(`${this.baseUrl}/api/collections`, {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${this.jwtToken}`
            }
        });

        if (res.status == 200) {
            return true;
        }

        if (await this.AuthRefresh()) {
            return true;
        }

        this.jwtToken = '';
        return false;
    }

    public async Login(username: string, password: string, superuser: boolean): Promise<boolean> {
        this.isSuperUser = superuser;
        const sub = superuser ? "_superusers" : "users";
        const res = await fetch(`${this.baseUrl}/api/collections/${sub}/auth-with-password`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                identity: username,
                password: password
            }),
        });

        if (res.status !== 200) {
            return false;
        }

        const result = await res.json() as LoginResponse;
        if (!result.token) {
            return false;
        }

        this.jwtToken = result.token;
        localStorage.setItem("JWT", result.token);
        return true;
    }

    public async AuthRefresh(): Promise<boolean> {
        const sub = this.isSuperUser ? "_superusers" : "users";
        const res = await fetch(`${this.baseUrl}/api/collections/${sub}/auth-refresh`, {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${this.jwtToken}`
            }
        });

        if (res.status !== 200) {
            return false;
        }

        const result = await res.json() as LoginResponse;
        this.jwtToken = result.token;
        localStorage.setItem("JWT", result.token);

        return true;
    }

    public async ListAllRepos(): Promise<RepoRecord[]> {
        const res = await fetch(`${this.baseUrl}/api/collections/repos/records`, {
            headers: {
                "Authorization": `Bearer ${this.jwtToken}`
            }
        });
        if (res.status == 401 || res.status == 403) {
            await goto("/login");
        }

        const items = await res.json() as CollectionsListResponse<RepoRecord>;
        return items.items;
    }
}
