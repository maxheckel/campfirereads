
export function capitalize(word) {
    return word
        .toLowerCase()
        .replace(/\w/, firstLetter => firstLetter.toUpperCase());
}

export function bookHref(book){
    let isbn = book.volumeInfo?.industryIdentifiers?.find((ii) => ii.type == 'ISBN_13')
    if (isbn){
        return '/book/'+isbn.identifier+'?v='+book.id;
    }
    return undefined;
}

export const decimalHash = string => {
    let sum = 0;
    for (let i = 0; i < string.length; i++)
        sum += (i + 1) * string.codePointAt(i) / (1 << 8)
    let result = sum % 1;
    if (result > 0.5) {
        return result / 2
    }
    return result;
}


export function imageUrl(book) {
    return book.volumeInfo.imageLinks?.thumbnail?.replace("edge=curl", "").replace("http://", "https://").replace('zoom=1', 'zoom=50')
}
