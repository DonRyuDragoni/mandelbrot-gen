type point =
    { x : int
    ; y : int
    };;

let next_point p cols =
    let new_x =
        succ p.x;

    and new_y =
        succ p.y;

    in if new_x >= cols then
        {x = 0; y = new_y}
    else
        {p with x = new_x};;

let prev_point p cols =
    let new_x =
        pred p.x;

    and new_y =
        pred p.y;

    in if new_x < 0 then
        {x = (pred cols); y = new_y}
    else
        {p with x = new_x};;

