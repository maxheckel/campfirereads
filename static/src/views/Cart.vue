<template>
  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 p-4">
    <Header class="ml-8" :text="'Your Cart'" :width-override="'w-10'" :icon-path="'/media/arrowhead.svg'"/>
    <div class="grid sm:grid-cols-[70%_30%] gap-8 p-4" v-if="cart.items.length > 0">
      <div>
        <div v-for="(item, index) in cart.items"
             class="grid md:grid-cols-[20%_80%] my-4 border-b-1 border-b-gray-200 border-b py-4 border-box">
          <a :href="bookHref(item.book)">
            <img :src='thumbnailUrl(item.book)' class="w-2/3 relative mx-auto ">
          </a>
          <div class="relative">
            <h2 class="text-2xl">{{ item.book.volumeInfo.title }}</h2>
            <div class="italic">By {{ item.book.volumeInfo.authors.join(', ').trim(', ') }}</div>
            <div v-if="!data.isbnToLoadingPrice[bookToISBN(item.book)]">
              {{ capitalize(item.listing.type) }} ${{ (item.listing.price_in_cents / 100).toFixed(2) }}
            </div>
            <div v-else>
              Loading Price...
            </div>
            <Button @click="removeFromCartAtIndex(index)"
                    class="text-sm my-4 font-normal ml-auto block top-0 md:absolute right-0 py-1 px-1 border-2  !hover:bg-red-200 border-red-200 text-np-dark-brown !bg-white"
                    :text="'Remove'"></Button>

          </div>

        </div>
      </div>

      <div>

        <div class="border border-gray-500 rounded-md h-auto p-4 border-dashed	bg-np-yellow-200">
          <b class="text-lg">Summary</b>

          <CartLineItem :value="''+cart.items.length" :label="'Items'"></CartLineItem>
          <template v-if="!loadingAnyPrices()">
            <CartLineItem :value="'$'+(subtotal()/100).toFixed(2)" :label="'Subtotal'"></CartLineItem>
            <CartLineItem :value="'$'+(smoke()).toFixed(2)" :label="'Smoke'"></CartLineItem>
            <CartLineItem :value="'TBD'" :label="'Shipping'"></CartLineItem>
            <CartLineItem :value="'$'+(total()/100).toFixed(2)" :label="'Total Before Shipping'"></CartLineItem>



            <Button @click="goToCheckout()" v-if="!data.loadingCheckout" class="w-full text-center mt-10"
                    :text="'Proceed to Checkout'"></Button>


          </template>
          <Loading class="relative mx-auto" v-if="data.loadingCheckout || loadingAnyPrices()"></Loading>
        </div>
        <Promo class="p-2 mt-2 bg-gray-200 text-np-dark-brown"/>
      </div>

    </div>
    <div v-if="cart.items.length == 0" class="text-4xl my-10 ml-16">
      You have no items in your cart
    </div>

  </div>

</template>

<script setup>

import Header from "../components/Header.vue";
import {bookToISBN, cart, removeFromCartAtIndex, removeISBNWithListingType, updatePrice} from "../store/cart.js";
import {bookHref, capitalize, thumbnailUrl} from "../services/utils.js";
import Button from "../components/Button.vue";
import CartLineItem from "../components/CartLineItem.vue";
import {onMounted, reactive} from "vue";
import Loading from "../components/icons/Loading.vue";
import {getSmoke} from "../store/cost.js";
import Promo from "../components/Promo.vue";

const data = reactive({
  loadingCheckout: false,
  loadingPrices: false,
  isbnToLoadingPrice: {}
})

function loadingAnyPrices() {
  for (let x = 0; x < Object.keys(data.isbnToLoadingPrice).length; x++) {
    if (data.isbnToLoadingPrice[Object.keys(data.isbnToLoadingPrice)[x]]) {
      return true
    }
  }
}

// Refresh the prices if they're > 1d old.  This prevents us from having skew when checking out.
onMounted(() => {
  var OneDay = new Date().getTime() - (24 * 60 * 60 * 1000)
  cart.items.forEach((item) => {
    if (item.listing.price_in_cents <= 0){

      removeISBNWithListingType(bookToISBN(item.book), item.listing.type)
      alert(item.book.volumeInfo.title + "is not available, it has been removed from your cart.")
      return
    }
    if (new Date(item.addedOn).getTime() < OneDay) {
      data.isbnToLoadingPrice[bookToISBN(item.book)] = true
      fetch(import.meta.env.VITE_API_HOST + "isbn/" + bookToISBN(item.book) + '/price')
          .then((response) => response.json())
          .then((resp) => {
            let found = false;
            resp.listings.forEach((l) => {
              if (l.type === item.listing.type) {
                found = true;
                item.addedOn = new Date();
                updatePrice(bookToISBN(item.book), item.listing.type, l.price_in_cents)
              }
            })

            if (!found) {
              removeISBNWithListingType(bookToISBN(item.book), item.listing.type)
              alert("Some items in your cart are no longer available, they've been automatically removed.")
            }

            data.isbnToLoadingPrice[bookToISBN(item.book)] = false

          });
    }

  })
})

function goToCheckout() {
  data.loadingCheckout = true
  fetch(import.meta.env.VITE_API_HOST + "checkout", {
    method: 'POST', // *GET, POST, PUT, DELETE, etc.
    headers: {
      'Content-Type': 'application/json'
      // 'Content-Type': 'application/x-www-form-urlencoded',
    },
    referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
    body: JSON.stringify(cart.items) // body data type must match "Content-Type" header
  })
      .then((response) => response.json())
      .then((resp) => {
        if (resp.url) {
          window.location = resp.url
          return
        }
        if (resp.type === "price_mismatch") {
          updatePrice(resp.data.isbn, resp.data.listingType, resp.data.actualPrice)
          alert(resp.error)
          data.loadingCheckout = false
          return
        }
        if (resp.type === "out_of_stock") {
          removeISBNWithListingType(resp.data.isbn, resp.data.listingType)
          alert(resp.error)
          data.loadingCheckout = false
          return
        }
        data.loadingCheckout = false
        alert("Something went wrong, please try again later")
      })
      .catch((err) => {
        alert(err)
        data.loadingCheckout = false
      });
}

function subtotal() {
  let total = 0;
  for (let x = 0; x < cart.items.length; x++) {
    total += cart.items[x].listing.price_in_cents;
  }
  return total;
}

function smoke() {
  return getSmoke().cost/100;
}

function total() {
  let total = 0;
  for (let x = 0; x < cart.items.length; x++) {
    total += cart.items[x].listing.price_in_cents;
  }
  total += getSmoke().cost
  return total;
}
</script>

<style scoped>

</style>