<script setup>
import {ref} from "vue";
import axios from "../utils/axios/axios";

let indexes = ref([])
let content = ref(null)
let page = ref({
  from: 0,
  size: 30,
})

const basicSearch = async () => {
  try {
    const response = await axios.post('/data/basic/search', {
          content: content.value,
          indexes: [],
          from: page.value.from,
          size: page.value.size,
        },
    );
    console.log(response.data);
  } catch (error) {
    console.error('Error fetching data:', error);
  }
}
</script>

<template>
  <div>
    <div class="search-container">
      <input type="text" placeholder="Search..." class="search-input" id="searchInput" v-model="content">
      <button class="search-button" @click="basicSearch">Search</button>
    </div>
  </div>
</template>

<style scoped>
.search-container {
  display: flex;
  align-items: center;
  width: 70%;
}

.search-input {
  flex: 1;
  margin: 1rem 1rem 1rem 0;
  border: 2px #9a9a9a solid;
  border-radius: 5px;
  height: 2.5rem;
  outline: none;
  font-size: 1rem;
}

.search-button {
  height: 2.8rem;
  border: 2px #9a9a9a solid;
  border-radius: 5px;
  background-color: #fff;
  cursor: pointer;
}

@media (prefers-color-scheme: dark) {
  body {
    background-color: #333;
    color: #fff;
  }

  .search-input {
    background-color: #444;
    color: #fff;
    border-color: #666;
  }

  .search-button {
    background-color: #555;
  }
}
</style>
