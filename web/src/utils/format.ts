import { DEFAULT_STEP_OPTIONS } from "./constants";

export function capitalizeWords(inputString: string | null) {
    if (!inputString) {
        return ""
    }
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

type InputObject = { [key: string]: any };

export function copyBody(inputObject: InputObject): InputObject {
    const processedObject: InputObject = {};

    for (const [key, value] of Object.entries(inputObject)) {

        if (key in DEFAULT_STEP_OPTIONS && typeof DEFAULT_STEP_OPTIONS[key] === 'function') {
            processedObject[key] = DEFAULT_STEP_OPTIONS[key]();
            continue
        }

        // Object and Number
        if (typeof value === 'number' ||
            typeof value === 'object' ||
            // Hexadecimal
            (typeof value === 'string' && value.startsWith("0x")) ||
            // Hexadecimal
            (typeof value === 'string' && /^[0-9a-fA-F]{24}$/.test(value)) ||
            // Mongodb Id
            (typeof value === 'string' && value.length === 24 && /^[0-9a-fA-F]{24}$/.test(value)) ||
            // Path
            (typeof value === 'string' && value.includes("/"))) {
            // Keep the key-value pair the same
            processedObject[key] = value;
        } else if (typeof value === 'string') {
            // Add "(copy)" to the end of the string
            processedObject[key] = `${value} (copy)`;
        }
    }

    return processedObject;
}

export function getRandomNumber(min: number, max: number): number {
    if (min > max) {
        [min, max] = [max, min];
    }

    const randomFraction = Math.random();
    return Math.floor(randomFraction * (max - min + 1)) + min;
}