<template>
  <div>
    <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-gray-300">Search</label>
    <div class="relative">
      <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
        <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor"
             viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
        </svg>
      </div>
      <input v-model="data.query" type="search" id="default-search"
             @keyup.enter="search"
             class="block p-4 pl-10 w-full bg-g text-sm text-gray-900 rounded-full border border-gray-300 focus:ring-blue-500 focus:border-blue-500 d dark:border-gray-600 dark:placeholder-gray-400 dark:text-np-dark-brown dark:focus:ring-blue-500 dark:focus:border-blue-500"
             :class="{'greenBG': greenBG}"
             placeholder="Search for any book, category, genre, or subject!" required>
      <button @click="search" type="submit"
              :class="{'greenBG': greenBG}"
              class="text-white absolute right-2.5 bottom-2.5 bg-np-green hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm px-4 py-2 dark:bg-np-green  dark:focus:ring-blue-800">
        Search
      </button>
    </div>
  </div>
</template>

<script setup>
import {reactive} from "vue";
import router from "../router";
import {header} from "../store/header.js";
const props = defineProps({
  greenBG: Boolean
})

const data = reactive({
  query: ''
})

function search() {
  window.location = '/search?query='+data.query;
}
</script>

<style scoped>
  input.greenBG{
    background-color: #4c5f4b;
    border: 1px solid #000;
    color: white;
  }
  button.greenBG{
    background-color: black;
  }
</style>