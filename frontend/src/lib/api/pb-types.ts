export interface CollectionsListResponse<T> {
    items: T[];
    page: number;
    perPage: number;
    totalItems: number;
    totalPages: number;
}

export interface LoginResponse {
    token: string,
    record: {
        collectionId: string;
        collectionName: string;
        id: string;
        email: string;
        emailVisibility: boolean;
        verified: boolean;
        created: string;
        updated: string;

    }
}