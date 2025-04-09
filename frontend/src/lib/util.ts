// Helper function to format dates
export function formatDate(dateString: string | Date | undefined): string {
    if (!dateString) return '-';
    const date = typeof dateString === 'string' ? new Date(dateString) : dateString;
    return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
}

export function formatBytes(bytes: number) {
    if (bytes === 0 || !bytes) {
        return '0 Bytes';
    }

    const k = 1024;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}