export const decoded = (b64: string) => {
    if (b64) {
        return window.atob(b64)
    }
    return null
}