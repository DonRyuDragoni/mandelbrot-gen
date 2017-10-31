open Complex;;
open Point;;

module Gr = Graphics;;

let remapf n (start1, stop1) (start2, stop2) =
    ((n -. start1)/.(stop1 -. start1)) *. (stop2 -. start2) +. start2;;

let remap ni (start1i, stop1i) (start2i, stop2i) =
    let n = float_of_int ni;
    and start1 = float_of_int start1i;
    and stop1  = float_of_int stop1i;
    and start2 = float_of_int start2i;
    and stop2  = float_of_int stop2i;
    in remapf n (start1, stop1) (start2, stop2);;

(*
 * Open a window with a given size.
 *)
let win size =
    Gr.open_graph
        (" " ^ (string_of_int size.x) ^ "x" ^ (string_of_int size.y));;

(*
 * Check if the corresponding complex number of the (x, y) point
 *
 *     z = x + y*i
 *
 * remains bounded when the function
 *
 *     f(z) = z^2 + c
 *
 * is applied to it multiple times
 *
 *     f(z), f(f(z)), f(f(f(z))), ...
 *
 * Returns the number of iterations.
 *)
let point_diverge max_iters x y =
    let rec div_r counter c z =
        if counter < max_iters && (norm2 z) < 4. then
            div_r (succ counter) c (add (mul z z) c)
        else
            counter

    and z = {re = x; im = y};

    in div_r 0 z z;;

(*
 * Generate the mandelbrot set.
 *
 * Each element of the list is the number of iterations its corresponding
 * complex number took in `point_diverge`.
 *)
let gen_set ?(max_iters = 63) size_x size_y =
    let rec divergence_lst acc p =
        let new_p = prev_point p size_x;

        in if p.y < 0 then
            acc
        else
            let x' =
                remapf
                    (float_of_int p.x)
                    (0.    , (float_of_int (pred size_x)))
                    (- 2.5 , 1.);
            and y' =
                remapf
                    (float_of_int p.y)
                    (0.  , (float_of_int (pred size_y)))
                    (-1. , 1.);

            in divergence_lst ((point_diverge max_iters x' y') :: acc) new_p;

    in divergence_lst [] {x = size_x; y = size_y};;

(*
 * Given a set with `cols` columns, draw it.
 *
 * This function assumes a window is opened.
 *)
let draw_set set cols =
    let set_color p =
        let b = int_of_float (remap p (0, 63) (0, 255));
        in Gr.set_color (Gr.rgb b b b);

    in let rec draw_points point = function
        | [] -> ()
        | p::ps ->
            set_color p;
            Gr.plot point.x point.y;
            draw_points (next_point point cols) ps;

    in draw_points {x = 0; y = 0} set;;

let size h =
    let w = h*16/9
    in { x = w
       ; y = h
       };;

(*
 * Generate and draw the set in a window.
 *)
let () =
    let win_size = size 600
    in
    win win_size;
    draw_set (gen_set win_size.x win_size.y) win_size.x;

    let rec evnt_loop finished =
        if (not finished) then
            evnt_loop
                (match Gr.read_key () with
                 | 'q' -> true
                 |  _  -> false)
    in
    evnt_loop false;;
