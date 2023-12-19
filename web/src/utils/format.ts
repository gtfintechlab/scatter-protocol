export function capitalizeWords(inputString: string) {
    return inputString.replace(/\b\w/g, (char: string) => char.toUpperCase());
}
