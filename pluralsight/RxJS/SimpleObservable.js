import { Observable } from "rxjs";
import { heroes } from "./data";

const subscribe = subscriber => {
    for (let hero of heroes) {
        if(hero["origin"] === "alien") {
            subscriber.error("Alien Found !!")
        }
        subscriber.next(hero)
    }
}

const heroesObservable$ = Observable.create(subscribe);

heroesObservable$.subscribe(value => console.log(value));
