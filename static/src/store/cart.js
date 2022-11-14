import {reactive} from "vue";

export const cart = reactive({
    items: [],
    justAdded: false,
    latestItem: {}
})

const cartKey = 'cart-1';

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


export function removeISBNWithListingType(isbn, type){
    for(let i = 0; i < cart.items.length; i++){
        if (cart.items[i].listing.type !== listingType){
            continue
        }
        let found = false
        // TODO: Gross
        for(let v = 0; v < cart.items[i].book.volumeInfo.industryIdentifiers.length; v++){
            if (cart.items[i].book.volumeInfo.industryIdentifiers[v].type === "ISBN_13" && cart.items[i].book.volumeInfo.industryIdentifiers[v].identifier === isbn){
                found = true
                break;
            }
        }
        if(!found){
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
        let found = false
        // TODO: Gross
        for(let v = 0; v < cart.items[i].book.volumeInfo.industryIdentifiers.length; v++){
            if (cart.items[i].book.volumeInfo.industryIdentifiers[v].type === "ISBN_13" && cart.items[i].book.volumeInfo.industryIdentifiers[v].identifier === isbn){
                found = true
                break;
            }
        }
        if(!found){
            continue
        }
        alert('here')
        alert(cart.items[i].listing.price_in_cents)
        cart.items[i].listing.price_in_cents = newPrice
        alert(cart.items[i].listing.price_in_cents)
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

