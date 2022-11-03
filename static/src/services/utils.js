
export function capitalize(word) {
    return word
        .toLowerCase()
        .replace(/\w/, firstLetter => firstLetter.toUpperCase());
}

export function bookHref(book){
    const isbn = book.volumeInfo.industryIdentifiers.find((ii) => ii.type == 'ISBN_13').identifier
    return '/book/'+isbn;
}