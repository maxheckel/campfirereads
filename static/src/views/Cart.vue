<template>
  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 p-4">
    <Header class="ml-8" :text="'Your Cart'" :width-override="'w-10'" :icon-path="'/media/arrowhead.svg'"/>
    <div class="grid sm:grid-cols-[70%_30%] gap-8 p-4" v-if="cart.items.length > 0">
      <div>
        <div v-for="(item, index) in cart.items" class="grid md:grid-cols-[20%_80%] my-4 border-b-1 border-b-gray-200 border-b py-4 border-box">
          <a :href="bookHref(item.book)">
            <img :src="imageUrl(item.book)" class="w-2/3 relative mx-auto ">
          </a>
          <div class="relative">
            <h2 class="text-2xl">{{item.book.volumeInfo.title}}</h2>
            <div class="italic">By {{item.book.volumeInfo.authors.join(', ').trim(', ')}}</div>
            <div>
              {{capitalize(item.listing.type)}} ${{(item.listing.price_in_cents+1000)/100}}
            </div>
            <Button  @click="removeFromCartAtIndex(index)" class="text-sm my-4 font-normal ml-auto block top-0 md:absolute right-0 py-1 px-1 border-2  !hover:bg-red-200 border-red-200 text-np-dark-brown !bg-white" :text="'Remove'"></Button>

          </div>

        </div>
      </div>

      <div >
        <div class="border border-gray-500 rounded-md h-auto p-4">
          <b class="text-lg">Summary</b>

          <CartLineItem :value="cart.items.length" :label="'Items'"></CartLineItem>
          <CartLineItem :value="'$'+(subtotal()/100)" :label="'Subtotal'"></CartLineItem>
          <CartLineItem :value="'$'+(smoke())" :label="'Smoke'"></CartLineItem>
          <CartLineItem :value="'$'+(total()/100)" :label="'Total'"></CartLineItem>
          <Button class="w-full text-center mt-10" :text="'Proceed to Checkout'"></Button>
        </div>

      </div>
    </div>
    <div v-if="cart.items.length == 0" class="text-4xl my-10 ml-16">
      You have no items in your cart
    </div>

  </div>

</template>

<script setup>

import Header from "../components/Header.vue";
import {cart, removeFromCartAtIndex} from "../store/cart.js";
import {capitalize} from "../services/utils.js";
import Button from "../components/Button.vue";
import CartLineItem from "../components/CartLineItem.vue";
import {bookHref, imageUrl} from "../services/utils.js";

function subtotal(){
  let total = 0;
  for(let x = 0; x < cart.items.length; x++){
    total += cart.items[x].listing.price_in_cents;
  }
  return total;
}

function smoke(){
  return cart.items.length * 10;
}

function total(){
  let total = 0;
  for(let x = 0; x < cart.items.length; x++){
    total += cart.items[x].listing.price_in_cents + 1000;
  }
  return total;
}
</script>

<style scoped>

</style>