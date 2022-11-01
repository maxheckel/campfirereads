<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
import BookList from "./components/BookList.vue";

</script>

<template class="min-h-screen">
  <header class="w-full p-4 h-20 bg-np-green  flex items-center rounded-bl-lg rounded-br-lg -mb-2 z-10 relative block  shadow-lg">
    <img src="/media/pine.svg" class="h-full mr-4 opacity-50">
    <span class="text-4xl text-np-yellow uppercase font-light inline-block">
      Campfire Reads
      <Div class="block text-lg normal-case text-center text-np-dark-brown" style="font-family: Freehand">
        Premium Smoked Books
      </Div>
    </span>
  </header>
  <router-view/>

</template>

<style scoped lang="scss">


</style>
