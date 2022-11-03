import {reactive} from "vue";

export const cart = reactive({
    items: [],
    justAdded: false,
    latestItem: {}
})

export function addToCart(book){
    cart.items.push(book)
    localStorage.setItem('cart', JSON.stringify(cart))
    cart.justAdded = true
    cart.latestItem = book
    // setTimeout(()=>{
    //     cart.justAdded = false
    // }, 5000)
}

export function removeFromCartAtIndex(index){
    if (confirm('Are you sure you would like to remove "'+cart.items[index].book.volumeInfo.title+'" from your cart?')){
        cart.items.splice(index, 1)
        localStorage.setItem('cart', JSON.stringify(cart))
    }
}



export function loadFromLS(){
    if(!localStorage.getItem('cart')){
        return;
    }
    const loaded = JSON.parse(localStorage.getItem('cart'))
    cart.items = loaded.items;
}

