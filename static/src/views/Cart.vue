<template>
  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 p-4">
    <Header :text="'Cart'" :icon-path="'/media/pine.svg'"/>
    <div class="grid grid-cols-[80%_20%] gap-8" v-if="cart.items.length > 0">
      <div>
        <div v-for="(item, index) in cart.items" class="grid grid-cols-[20%_80%] my-4 border-b-1 border-b-gray-200 border-b py-4 border-box">
          <div>
            <img :src="item.book.volumeInfo.imageLinks.thumbnail" class="w-2/3 relative mx-auto ">
          </div>
          <div>
            <h2 class="text-2xl">{{item.book.volumeInfo.title}}</h2>
            <div class="italic">By {{item.book.volumeInfo.authors.join(', ').trim(', ')}}</div>
            <div>
              {{capitalize(item.listing.type)}} ${{(item.listing.price_in_cents+1000)/100}}
            </div>
            <Button  @click="removeFromCartAtIndex(index)" class="text-sm my-4 font-normal bg-red-400 text-white" :text="'Remove'"></Button>

          </div>

        </div>
      </div>

      <div class="border p-4">
        <b class="text-lg">Summary</b>
        <div class="flex my-4">
          <b>Items</b>
          <div class="ml-auto">
            {{cart.items.length}}
          </div>
        </div>
        <div class="flex my-4">
          <b>Total</b>
          <div class="ml-auto">
            ${{total()/100}}
          </div>
        </div>
      </div>
    </div>
    <div v-if="cart.items.length == 0">
      You have no items in your cart
    </div>

  </div>

</template>

<script setup>

import Header from "../components/Header.vue";
import {cart, removeFromCartAtIndex} from "../store/cart.js";
import {capitalize} from "../services/utils.js";
import Button from "../components/Button.vue";

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