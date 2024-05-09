export const sanitizeData = (obj: any): void => {
    for (let key in obj) {
        if (typeof obj[key] === 'string') {
            obj[key] = obj[key].trim()
        } else if (typeof obj[key] === 'object') {
            sanitizeData(obj[key])
        }
    }
}