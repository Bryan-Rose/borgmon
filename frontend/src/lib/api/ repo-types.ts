/*
* DB repo
*/

export interface RepoRecord {
    id: string;
    name: string;
    path: string;
    created: string;
    borgUpdated: string;
    lastBackup: string;
    updated: string;
    borgList: RepoList;
    borgInfo: RepoInfo;
}


/*
* borg --list
*/

export interface RepoList {
    archives: ListArchive[];
    encryption: ListEncryption;
    repository: ListRepository;
}

export interface ListArchive {
    archive: string;
    barchive: string;
    id: string;
    name: string;
    start: Date;
    time: Date;
}

export interface ListEncryption {
    mode: string;
}

export interface ListRepository {
    id: string;
    last_modified: Date;
    location: string;
}


/*
* borg --info
*/

export interface RepoInfo {
    cache: InfoCache;
    encryption: InfoEncryption;
    repository: InfoRepository;
    security_dir: string;
}

export interface InfoCache {
    path: string;
    stats: InfoStats;
}

export interface InfoStats {
    total_chunks: number;
    total_csize: number;
    total_size: number;
    total_unique_chunks: number;
    unique_csize: number;
    unique_size: number;
}

export interface InfoEncryption {
    mode: string;
}

export interface InfoRepository {
    id: string;
    last_modified: Date;
    location: string;
}
