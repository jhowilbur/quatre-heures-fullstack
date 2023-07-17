<template>
  <div class="list row">
    <div class="col-md-6">
      <h4>Events List</h4>
      <ul class="list-group">
        <li
          class="list-group-item"
          :class="{ active: index == currentIndex }"
          v-for="(event, index) in events"
          :key="index"
          @click="setActiveEvent(event, index)"
        >
          {{ event.title }}
        </li>
      </ul>
    </div>
    <div class="col-md-6">
      <div v-if="currentEvent.id">
        <h4>Event</h4>
        <div>
          <label><strong>Title:</strong></label> {{ currentEvent.title }}
        </div>
        <div>
          <label><strong>Description:</strong></label>
          {{ currentEvent.description }}
        </div>
        <div>
          <label><strong>Initial Date:</strong></label>
          {{ currentEvent.initial_date }}
        </div>
        <div>
          <label><strong>Final Date:</strong></label>
          {{ currentEvent.final_date }}
        </div>

        <router-link
          :to="'/events/' + currentEvent.id"
          class="badge badge-warning"
        >
          Edit
        </router-link>
      </div>
      <div v-else>
        <br />
        <p>Please click on an Event...</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import EventDataService from "@/services/EventDataService";
import Event from "@/types/Event";
import ResponseData from "@/types/ResponseData";

export default defineComponent({
  name: "events-list",
  data() {
    return {
      events: [] as Event[],
      currentEvent: {} as Event,
      currentIndex: -1,
      title: "",
    };
  },
  methods: {
    retrieveEvents() {
      EventDataService.getAll()
        .then((response: ResponseData) => {
          this.events = response.data;
          console.log(response.data);
        })
        .catch((e: Error) => {
          console.log(e);
        });
    },

    setActiveEvent(event: Event, index = -1) {
      this.currentEvent = event;
      this.currentIndex = index;
    },
  },
  mounted() {
    this.retrieveEvents();
  },
});
</script>

<style>
.list {
  text-align: left;
  max-width: 750px;
  margin: auto;
}
</style>
