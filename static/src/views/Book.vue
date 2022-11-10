<template>
  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 mt-20 p-4">
    <div class="grid md:grid-cols-[30%_70%] gap-16">
      <div>
        <img v-if="!data.loadingBook" :src="imageUrl(data.book)"
             class="shadow shadow-xlg w-1/2 relative mx-auto md:w-full">
        <ShimmerBox v-else class="w-full h-[400px]"></ShimmerBox>
      </div>
      <div>
        <template v-if="!data.loadingBook">
          <h1 class="text-4xl font-bold">{{ data.book?.volumeInfo?.title }}</h1>
          <b>By {{ data.book?.volumeInfo?.authors.join(', ').trim(', ') }}</b>
          <div v-if="data.loadingPrice">
            Loading price...
          </div>
          <template v-else>
            <div v-if="getListings().length > 1" class="">
              ${{ lowestPrice() / 100 }} - ${{ highestPrice() / 100 }}
            </div>
            <div v-else-if="getListings().length > 0" class="text-xl my-2">
              {{ capitalize(getListings()[0].type) }} ${{ (getListings()[0].price_in_cents + 1000) / 100 }}
            </div>
          </template>

          <h2 class="my-8">{{ description() }}</h2>
          <template v-if="!data.descriptionIsSmall">
            <span class="block cursor-pointer text-gray-500" @click="data.showingFullDescription = true"
                  v-if="!data.showingFullDescription">Show more</span>
            <span class="block cursor-pointer text-gray-500" @click="data.showingFullDescription = false"
                  v-if="data.showingFullDescription">Show less</span>
          </template>

          <select v-model="data.selectedListing" class=" block p-2 px-6 border rounded-md rounded mt-4 text-lg"
                  v-if="getListings().length > 1 && !data.loading">
            <option :value="i" :selected="data.selectedListing === i" v-for="(listing, i) in getListings()">
              {{ capitalize(listing.type) }} ${{ (listing.price_in_cents + 1000) / 100 }}
            </option>
          </select>
          <div class="block mt-4" v-if="isInCart()">
            <router-link :to="{name: 'cart'}" class="rounded-full px-4  bg-green-200 relative w-auto inline-block">
              <div class="flex items-center py-1 text-sm">
              <span class="material-symbols-outlined text-sm inline-block mr-1">
                local_mall
              </span>

                In Cart
              </div>

            </router-link>
          </div>
          <Button v-if="!data.loadingPrice" @click="formatAndAddToCart()" class="my-4 block"
                  :text="'Add to Cart'"></Button>
          <Button v-if="data.loadingPrice" class="bg-gray-200 text-gray-400 my-4 block"
                  :text="'Loading Price'"></Button>
        </template>
        <template v-else>
          <shimmer-box class="w-40 h-8 rounded rounded-full"/>
          <shimmer-box class="w-full rounded rounded-full h-4 mt-4" v-for="i in new Array(3)"></shimmer-box>
        </template>

      </div>
    </div>
  </div>
</template>

<script setup>

import {useRoute} from 'vue-router';
import {onMounted, reactive} from "vue";
import ShimmerBox from "../components/ShimmerBox.vue";
import Button from "../components/Button.vue";
import {addToCart, cart} from "../store/cart.js";
import {capitalize, imageUrl} from "../services/utils.js";

const route = useRoute();
const isbn = route.params.isbn

const data = reactive({
  book: {},
  prices: [],
  loadingBook: true,
  loadingPrice: true,
  showingFullDescription: false,
  descriptionIsSmall: false,
  selectedListing: 1
})

function formatAndAddToCart() {
  let listing = getListings()[data.selectedListing];
  if (listing === undefined) {
    listing = getListings()[0]
  }
  addToCart({
    book: data.book,
    listing: listing
  })


  document.body.scrollTop = document.documentElement.scrollTop = 0;

}

function isInCart() {

  for (let x = 0; x < cart.items.length; x++) {
    if (cart.items[x].book.id === data.book.id) {
      return true
    }
  }
  return false;
}


function getListings() {
  return data.prices.filter((l) => l.price_in_cents > 0)
}

function lowestPrice() {
  return getListings().sort((a, b) => a.price_in_cents - b.price_in_cents)[0].price_in_cents + 1000
}

function highestPrice() {
  return getListings().sort((a, b) => b.price_in_cents - a.price_in_cents)[0].price_in_cents + 1000
}

function description() {
  if (data.book.volumeInfo.description.split(' ').length < 50) {
    data.descriptionIsSmall = true;
    return data.book.volumeInfo.description;
  }
  if (data.showingFullDescription) {
    return data.book.volumeInfo.description;
  }
  return data.book.volumeInfo.description.split(' ').slice(0, 50).join(' ') + '...'
}

onMounted(() => {

  fetch(import.meta.env.VITE_API_HOST + "isbn/" + isbn)
      .then((response) => response.json())
      .then((resp) => {
        data.book = resp.book;
        data.loadingBook = false;
        if (data.book == null) {
          window.location.href = '/'
        }
      });
  fetch(import.meta.env.VITE_API_HOST + "isbn/" + isbn + '/price')
      .then((response) => response.json())
      .then((resp) => {
        data.prices = resp.listings;
        data.loadingPrice = false;
        if (data.prices.length == 0) {
          window.location.href = '/'
        }
      });
})
</script>

<style scoped>

</style>