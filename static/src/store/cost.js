import {reactive} from "vue";

const smoke = reactive({
    cost: 0,
    inFlight: false
})


export function getSmoke() {
    if (smoke.cost <= 0 && !smoke.inFlight) {
        smoke.inFlight = true
        fetch(import.meta.env.VITE_API_HOST + "cost")
            .then((response) => response.text()).then((response) => {
            smoke.cost = parseInt(response)
        })
    }
    return smoke
}