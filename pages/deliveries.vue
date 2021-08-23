<template>
  <div>
    <div v-for="row in routes" :key="row.id" :class="['maptag-card', getColor(row.status)]">
      <CardRoute title="Ruta" :data="row.id" />
      <CardRoute title="Conductor" :data="row.driver_name" />
      <CardRoute title="Creada" :data="row.created_at" />
      <CardRoute title="Inicio" :data="row.started_at" />
      <CardRoute title="Terminó" :data="row.ended_at" />
      <CardRoute title="Entregas" :data="row.deliveries" />
      <CardRoute title="" :data="getTranslate(row.status)" :classObject="getColor(row.status)" />
    </div>
  </div>
</template>

<script lang="js">
import Vue from 'vue';
import { MESES } from '@/constants/'

export default Vue.extend({
  layout: 'default',
  data() {
    return {
      routes: [],
      connection: null
    }
  },
  async fetch() {
    let result = await fetch(
        'https://798c19a6-3f80-43ae-a6a3-4964fea68347.mock.pstmn.io/routes'
    ).then(res => res.json());
    this.routes = result.routes.map(e => {
      let created_at = new Date('2011-04-11T10:20:30Z');
      e.created_at = created_at.getDay() + " de " + MESES[created_at.getMonth()] + " del " + created_at.getFullYear();
      e.started_at = created_at.getHours() + ":" + created_at.getMinutes();
      return e;
    });
    var self = this;
    this.connection = new WebSocket('ws://localhost:8080/')
    this.connection.onopen = function() {
      console.log("Connect to WebSocket")
    };
    this.connection.onmessage = function(event) {
      let route = JSON.parse(event.data);
      self.refresh(route);
    };
  },
  methods: {
    refresh(route) {
      let foundIndex = this.routes.findIndex( e => e.id == route.route_id);
      let new_router = this.routes[foundIndex];
      new_router.status = route.status;
      if (route.status == "onroute") {
        new_router.deliveries = new_router.deliveries + 1;
      } else {
        let completed_at = new Date( route.completed_at);
        new_router.ended_at = completed_at.getHours() + ":" + completed_at.getMinutes();
      }
      this.routes[foundIndex] = new_router;
    },
    getColor(status) {
      return {
        'blue': status == "pending" || status == "onroute",
        'green': status == "completed"
      }
    },
    getTranslate(status) {
      if (status == "pending") {
        return "Pendiente"
      }
      if (status == "onroute") {
        return "En ruta"
      }
      if (status == "completed") {
        return "Completada"
      }
    }
  }
})
</script>


<style>
.maptag-card {
  margin: 6px;
  padding: 6px;
  background-color: #ffffff;
  -webkit-border-radius: 3px;
  -moz-border-radius: 3px;
  border-radius: 3px;
}
.maptag-card.blue {
  border-left-style: solid;
  border-left-width: 4px;
  border-left-color: #3b8bff;
}
.maptag-card.green {
  border-left-style: solid;
  border-left-width: 4px;
  border-left-color: #52ba5d;
}
</style>
