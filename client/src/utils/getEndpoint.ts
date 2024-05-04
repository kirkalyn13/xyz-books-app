export const getEndpoint = (resource: string, id: string, path: string = ""): string => {
    return URL + "/api/v1/" + resource + "/" + path + (id ? "/" + id : "")
}