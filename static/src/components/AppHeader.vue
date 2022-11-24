<template>
  <header
      class="w-full p-4 h-20 bg-np-green  flex items-center rounded-bl-lg rounded-br-lg -mb-2 z-10 relative block  shadow-lg">
    <img src="/media/pine.svg" class="h-full mr-4 opacity-50">
    <a href="/" class="sm:text-4xl text-lg font-bold sm:font-normal text-np-yellow uppercase font-light inline-block">
      Campfire Reads
      <div class="block  normal-case text-sm sm:text-lg text-center text-np-dark-brown" style="font-family: Freehand">
        Premium Smoked Books
      </div>
    </a>
    <div class="ml-auto max-w-xl lg:w-full flex items-center">
      <Search :green-b-g="true" class="w-full hidden lg:block inline-block float-right mr-8"/>
      <span @click="header.showingMobileSearch = true" class="cursor-pointer material-symbols-outlined lg:hidden text-4xl font-bold mr-4">
      search
      </span>

      <a href="/cart" class="cursor-pointer inline-block ml-auto">
        <div class="absolute top-0 right-0">
          <div class="bg-np-yellow m-2 text-sm px-1 rounded-full rounded">
            {{ cart.items.length }}
          </div>

        </div>

        <span class="material-symbols-outlined text-3xl text-np-dark-brown font-bold">
      local_mall
      </span>
      </a>
    </div>

    <div v-if="cart.justAdded"
         class="absolute top-20 right-2 z-10 bg-white text-np-dark-brown border border-2 border-np-dark-brown w-2/3 md:w-80 shadow-xl p-4 -mt-4 rounded-md z-20">
      <div @click="cart.justAdded = false;"
           class="absolute top-2 right-4 text-xl font-bold  rounded-full cursor-pointer">
        x
      </div>
      <b>Just Added</b>
      <div class="grid grid-cols-[20%_80%] gap-2 my-2">
        <div>
          <img class="shadow-xl block relative" :src="cart.latestItem.book?.volumeInfo.imageLinks.thumbnail">
        </div>
        <div>
          <span class="font-bold block">{{ cart.latestItem.book?.volumeInfo.title }}</span>
          <span class="italic block">by {{ cart.latestItem.book?.volumeInfo.authors.join(', ').trim(', ') }}</span>
          {{ capitalize(cart.latestItem.listing.type) }} ${{ ((cart.latestItem.listing.price_in_cents + getSmoke().cost) / 100).toFixed(2) }}
          <a href="/cart">
            <Button
                class="text-sm ml-auto float-right bg-transparent text-black border-black border border-2 hover:bg-black hover:text-white mt-4 mr-4"
                :text="'Go to Cart'"></Button>
          </a>
        </div>

      </div>

    </div>
  </header>
  <div v-if="header.showingMobileSearch" class="t-20 flex lg:hidden w-full mt-4 absolute z-20 items-center">
    <Search class="w-full" />
    <div class="mx-4">
      <Button @click="header.showingMobileSearch = false" class="!bg-red-200 !text-sm text-gray-500 !px-2 !py-1 !rounded-full" :text="'X'" />
    </div>

  </div>
</template>

<script setup>

import {cart} from "../store/cart.js";

import {capitalize} from "../services/utils.js";
import Button from "../components/Button.vue";
import Search from "../components/Search.vue";
import {header} from "../store/header.js";
import {getSmoke} from "../store/cost.js";

</script>

<style scoped>

</style>