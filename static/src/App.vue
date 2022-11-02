<script setup>
import { RouterLink, RouterView } from 'vue-router'
import {cart} from "./store/cart.js";
import {onMounted} from "vue";
import {loadFromLS} from "./store/cart.js";
import Book from "./components/Book.vue";
onMounted(() => {
  loadFromLS()
})
function capitalize(word) {
  return word
      .toLowerCase()
      .replace(/\w/, firstLetter => firstLetter.toUpperCase());
}
</script>

<template class="min-h-screen">
  <header class="w-full p-4 h-20 bg-np-green  flex items-center rounded-bl-lg rounded-br-lg -mb-2 z-10 relative block  shadow-lg">
    <img src="/media/pine.svg" class="h-full mr-4 opacity-50">
    <a href="/" class="sm:text-4xl text-lg font-bold sm:font-normal text-np-yellow uppercase font-light inline-block">
      Campfire Reads
      <div class="block  normal-case text-sm sm:text-lg text-center text-np-dark-brown" style="font-family: Freehand">
        Premium Smoked Books
      </div>
    </a>
    <a href="/cart" class="ml-auto cursor-pointer">
      <div class="absolute top-0 right-0">
        <div class="bg-np-yellow m-2 text-sm px-1 rounded-full rounded">
          {{cart.items.length}}
        </div>

      </div>

      <span class="material-symbols-outlined text-3xl text-np-dark-brown font-bold">
      local_mall
      </span>
    </a>
    <div v-if="cart.justAdded" class="absolute top-20 right-2 z-10 bg-np-dark-brown text-np-yellow w-2/3 shadow-xl p-4 -mt-4 rounded-md z-20">
      <div @click="cart.justAdded = false;" class="absolute top-2 right-4 text-xl font-bold  rounded-full cursor-pointer">
        x
      </div>
      <b>Just Added</b>
      <div class="grid grid-cols-[20%_80%] gap-2 my-2">
        <div>
          <img class="shadow-xl block relative" :src="cart.latestItem.book?.volumeInfo.imageLinks.thumbnail">
        </div>
        <div>
          <span class="font-bold block">{{cart.latestItem.book?.volumeInfo.title}}</span>
          <span class="italic block">by {{cart.latestItem.book?.volumeInfo.authors.join(', ').trim(', ')}}</span>
          {{capitalize(cart.latestItem.listing.type)}} ${{(cart.latestItem.listing.price_in_cents+1000)/100}}
        </div>

      </div>

    </div>
  </header>
  <router-view/>

</template>

<style scoped lang="scss">


</style>
