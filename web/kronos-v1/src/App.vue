<script setup>
import {ref} from "vue";
import Search from "./views/Search.vue";
import Content from "./views/Content.vue";
import Settings from "./views/Settings.vue";
import axios from "./utils/axios/axios";

let indexes = ref([])

const selected = ref()

const handleClick = (indexName) => {
  selected.value = indexName
}

const fetchData = async () => {
  try {
    const response = await axios.get('/script/get/extensions');
    indexes.value = response.data;
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

fetchData();

</script>

<template>
  <div class="container">
    <div class="sidebar">
      <div class="option" @click="handleClick('Search')">
        Search
      </div>

      <hr/>

      <div v-if="indexes.length !== 0">
        <div v-for="item in indexes">
          <div class="option" @click="handleClick(item.name)">
            {{ item.name }}
          </div>
        </div>
        <hr/>
      </div>

      <div class="option" @click="handleClick('Settings')">
        Settings
      </div>
    </div>

    <div class="content">
      <div v-if="selected === 'Settings'">
        <Settings></Settings>
      </div>
      <div v-else class="content">
        <Search></Search>
        <Content :selectedPlugin="selected"></Content>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  flex-direction: column;
}

.content {
  flex: 1;
  height: 100vh;
  overflow-y: auto;
  margin-left: var(--sidebar-width);
}
</style>
