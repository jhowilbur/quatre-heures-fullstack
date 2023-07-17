<template>
  <div id="app">
    <nav class="navbar navbar-expand navbar-dark bg-dark">
      <router-link to="/" class="navbar-brand">Wilbur - Kumojin</router-link>
      <div class="navbar-nav mr-auto">
        <li class="nav-item">
          <router-link to="/events" class="nav-link">Events</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/add" class="nav-link">Add</router-link>
        </li>
      </div>
      <img :src="require('@/assets/kumojin.svg')" style="width: 150px" alt="" />
    </nav>

    <div class="container mt-3">
      <router-view />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
});
</script>
