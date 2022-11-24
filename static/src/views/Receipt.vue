<template>
  <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 p-4">
    <Header v-if="!data.loading" class="" :text="'Your Order On '+orderedOn()" :width-override="'w-10'" :icon-path="'/media/arrowhead.svg'"/>
    <span class="block ml-4" v-if="!data.loading">Order #{{data.receipt.orderID}}</span>
    <Loading class="relative mx-auto" v-if="data.loading"></Loading>

    <div class="grid sm:grid-cols-[70%_30%] gap-8 p-4" v-if="!data.loading">

      <div>
        <div class="grid sm:grid-cols-2 gap-8 p-4 border-1 border border-dashed	 border-black rounded-lg" v-if=" !data.loading">
          <div>
            <Address :address="data.receipt.shipping" :type="'Shipping Address'"/>
          </div>
          <div>
            <Address :address="data.receipt.billing" :type="'Billing Address'"/>
          </div>
        </div>
        <div v-for="(item) in data.receipt.books"
             class="grid md:grid-cols-[20%_80%] mb-4 border-b-1 border-b-gray-200 border-b py-4 border-box">
          <a :href="bookHref(item.book)">
            <img :src="imageUrl(item.book)" class="w-2/3 relative mx-auto ">
          </a>
          <div class="relative">
            <h2 class="text-2xl">{{ item.book.volumeInfo.title }}</h2>
            <div class="italic">By {{ item.book.volumeInfo.authors.join(', ').trim(', ') }}</div>
            <div>
              {{ capitalize(item.listing.type) }} ${{ ((item.listing.price_in_cents) / 100).toFixed(2) }}
            </div>

          </div>

        </div>
        <Support/>
      </div>

      <div>
        <div class="border border-gray-500 rounded-md h-auto p-4 border-dashed bg-np-yellow-200	">
          <b class="text-lg">Summary</b>
          <CartLineItem :value="capitalize(data.receipt.paymentType)" :label="'Payment Type'"></CartLineItem>
          <div class="sm:text-right text-np-dark-brown italic" v-if="data.receipt.paymentIdentifier">{{data.receipt.paymentIdentifier}}</div>
          <CartLineItem :value="''+data.receipt.books.length" :label="'Items'"></CartLineItem>
          <CartLineItem :value="'$'+(subtotal()/100).toFixed(2)" :label="'Subtotal'"></CartLineItem>
          <CartLineItem :value="'$'+(smoke()).toFixed(2)" :label="'Smoke'"></CartLineItem>
          <CartLineItem :value="'$'+(total()/100).toFixed(2)" :label="'Total'"></CartLineItem>


        </div>

      </div>
    </div>

  </div>
</template>

<script setup>
import {onMounted, reactive} from "vue";
import {useRoute} from "vue-router/dist/vue-router";
import {bookHref, capitalize, imageUrl} from "../services/utils";
import {clear} from "../store/cart";
import Loading from "../components/icons/Loading.vue";
import CartLineItem from "../components/CartLineItem.vue";
import Header from "../components/Header.vue";
import Address from "../components/Address.vue";
import Support from "../components/Support.vue";

const route = useRoute();
const id = route.params.id

const data = reactive({
  receipt: {},
  loading: true
})

onMounted(() => {

  fetch(import.meta.env.VITE_API_HOST + "receipt/" + id)
      .then((response) => response.json())
      .then((resp) => {
        var OneMinute = new Date().getTime() - (60 * getSmoke().cost)
        data.loading = false;
        data.receipt = resp
        if (new Date(data.receipt.orderedOn).getTime() > OneMinute){
          clear()
        }
      });
})


function subtotal() {
  let total = 0;
  data.receipt.books.forEach((b) => total+=b.listing.price_in_cents)
  return total - (data.receipt.books.length*getSmoke().cost);
}

function smoke() {
  return data.receipt.books.length * 10;
}

function orderedOn(){
  return new Date(data.receipt.orderedOn).toLocaleDateString()
}

function total() {
  return data.receipt.totalInCents
}

</script>

<style scoped>

</style>