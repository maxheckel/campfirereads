import {reactive} from "vue";

export const cart = reactive({
    items: [],
    justAdded: false,
    latestItem: {},
})

const cartKey = 'cart-1';

export function clear(){
    cart.items = []
    cart.justAdded = false;
    cart.latestItem = {};
    localStorage.setItem(cartKey, JSON.stringify(cart))
}

export function addToCart(book){
    cart.items.push(book)
    localStorage.setItem(cartKey, JSON.stringify(cart))
    cart.justAdded = true
    cart.latestItem = book

    // setTimeout(()=>{
    //     cart.justAdded = false
    // }, 5000)
}

export function removeFromCartAtIndex(index){
    if (confirm('Are you sure you would like to remove "'+cart.items[index].book.volumeInfo.title+'" from your cart?')){
        cart.items.splice(index, 1)
        localStorage.setItem(cartKey, JSON.stringify(cart))
    }
}

export function bookToISBN(book){
    for(let v = 0; v < book.volumeInfo.industryIdentifiers.length; v++){
        if (book.volumeInfo.industryIdentifiers[v].type === "ISBN_13"){
            return book.volumeInfo.industryIdentifiers[v].identifier
        }
    }
    return ""
}

export function removeISBNWithListingType(isbn, type){
    for(let i = 0; i < cart.items.length; i++){
        if (cart.items[i].listing.type !== type){
            continue
        }
        if (isbn !== bookToISBN(cart.items[i].book)){
            continue
        }

        cart.items.splice(i, 1)
    }
    localStorage.setItem(cartKey, JSON.stringify(cart))
}

export function updatePrice(isbn, listingType, newPrice){
    for(let i = 0; i < cart.items.length; i++){
        if (cart.items[i].listing.type !== listingType){
            continue
        }
        if (isbn !== bookToISBN(cart.items[i].book)){
            continue
        }
        cart.items[i].listing.price_in_cents = newPrice

    }
    localStorage.setItem(cartKey, JSON.stringify(cart))
}


export function loadFromLS(){
    if(!localStorage.getItem(cartKey)){
        return;
    }
    const loaded = JSON.parse(localStorage.getItem(cartKey))
    cart.items = loaded.items;
}

