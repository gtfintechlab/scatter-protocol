export function capitalizeWords(inputString: string) {
    return inputString.replace(/\b\w/g, (char: string) => char.toUpperCase());
}

export function enumToDictionary(enumObject: { [x: string]: string; }): Record<string, string> {
    const dictionary: Record<string, string> = {};

    for (const key in enumObject) {
        if (Object.prototype.hasOwnProperty.call(enumObject, key)) {
            dictionary[enumObject[key]] = enumObject[key];
        }
    }

    return dictionary;
}