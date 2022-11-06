<template>
  <div>
    <div class="text-center p-8 bg-center bg-cover	" :style="{'background-image': 'url('+data.bgImage+')'}">
      <div class="z-10 w-full  relative left-0 p-8">
        <div class="trapezoid text-np-yellow p-10 text-center w-full lg:w-2/3 relative">

          <Header :invert="true" class="text-white z-10 block mx-auto justify-center relative" :width-override="'w-10'"
                  :text="displayName()" :icon-path="'/media/arrowhead.svg'"></Header>
          <span class="text-np-yellow block relative z-10 md:text-2xl tracking-widest md:mt-4"
                style="font-family: Freehand">
          Best Sellers
        </span>
        </div>

      </div>
    </div>
    <div class="max-w-7xl mx-auto sm:px-6 lg:px-8 md:py-8 pt-8 p-4">
      <div class="grid grid-cols-2 sm:grid-cols-4 lg:grid-cols-5 gap-8">
        <Book v-if="!data.loading" v-for="book in data.list?.books" :book="book"></Book>
        <ShimmerBox v-if="data.loading === true" v-for="i in new Array(15)"
                    class="w-full h-[250px] bg-gray-200"></ShimmerBox>

      </div>
    </div>
  </div>
</template>

<script setup>
import {useRoute} from 'vue-router';
import {reactive} from "vue";
import Header from "../components/Header.vue";
import Book from "../components/Book.vue";
import ShimmerBox from "../components/ShimmerBox.vue";
import {decimalHash} from "../services/utils.js";
const route = useRoute();
const category = route.params.category
const data = reactive({
  list: {},
  loading: true,
  bgImage: ""
})

const displayName = () => {
  return data.list?.list?.display_name ? data.list?.list?.display_name : category.replace("-", " ")
}



fetch(import.meta.env.VITE_API_HOST + "/category/" + category)
    .then((response) => response.json())
    .then((resp) => {
      data.list = resp;
      data.bgImage = "/media/covers/" + (decimalHash(data.list.list.display_name) + "").split('.')[1].slice(0, 2) + ".jpg"
      data.loading = false;
    });
</script>

<style lang="scss">

.trapezoid {
  display: inline-block;

  position: relative;
}

.trapezoid::before,
.trapezoid::after {
  content: "";
  position: absolute;
  top: 0;
  border: 4px solid #F0E9CF;
  bottom: 0;
  width: 60%;
  background: #221B1C;
  transform-origin: top;
  box-sizing: border-box;
  z-index: 0;
}

.trapezoid::before {
  border-radius: 30px 0 0 30px;
  border-right: 0;
  left: 0;
  transform: skew(8deg);
}

.trapezoid::after {
  border-radius: 0 30px 30px 0;
  border-left: 0;
  right: 0;
  transform: skew(-8deg);
}


</style>