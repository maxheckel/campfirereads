<script setup>
import {RouterView} from 'vue-router'
import {cart, loadFromLS} from "./store/cart.js";
import {onMounted} from "vue";
import AppHeader from "./components/AppHeader.vue";

onMounted(() => {
  loadFromLS()
})
</script>

<template class="min-h-screen">
  <AppHeader/>
  <router-view/>
  <div class="footer-head w-full h-[100px] bg-np-dark-brown mt-20"></div>
  <footer class="w-full p-4 bg-np-dark-brown ">
    <div class="text-np-yellow  sm:flex items-center -mb-2 z-10 relative block  shadow-lg">
      <div class="sm:w-1/2 md:w-1/3 -mt-10">
        <h3 class="font-bold uppercase text-2xl">Campfire Reads</h3>
        <div class="text-np-yellow-200">
          Campfire Reads provides a marketplace for a large amount of books shipped directly to your house.
          <div class="text-xs my-4">
            <b>Smoked Book</b> (sməʊkt bo͝ok): A book that has been imbued with the smell of campfire. Instant nostalgia
            at your fingertips without leaving the hhouse.
          </div>
        </div>
      </div>

      <div class="ml-auto sm:text-right">
        <b>Links</b>
        <a class="block" href="/about">About Us</a>
        <a class="block" href="/contact">Contact</a>
        <a class="block" href="/contact">Terms of Service</a>
        <a class="block" href="/contact">Privacy Policy</a>

      </div>
    </div>
    <div class="text-center my-10 block">
      <span class="material-symbols-outlined text-8xl font-light text-np-yellow opacity-30">
      local_fire_department
      </span>
    </div>
  </footer>
</template>

<style scoped lang="scss">
.footer-head {
  transform: skewY(2deg);
  margin-bottom: -40px;

}

</style>
