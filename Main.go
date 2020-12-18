package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var x_arr = 1000
var y_arr = 1000
var Node_x [Nnode]int
var Node_y [Nnode]int
var dis [Nnode][Nnode]float64
var pij [Nnode][Nnode]float64
var passij [Nnode][Nnode]float64
const Nnode = 10

func main(){
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<Nnode; i++	{
		Node_x[i] = rand.Intn(x_arr)
		Node_y[i] = rand.Intn(y_arr)
	}
	for i:=0; i<Nnode; i++	{
		for j:=0; j<Nnode; j++ {
			var Delta_x = Node_x[i] - Node_x[j]
			var Delta_y = Node_y[i] - Node_y[j]
			dis[i][j] = math.Sqrt(float64(Delta_x*Delta_x + Delta_y*Delta_y))
			pij[i][j] = 1
			passij[i][j] = 0
		}
	}

	map_distance := make(map[int]float64)
	map_path := make(map[int][]int)
	var slice_ants_arrived []int
	var dis_ants []float64
	for i:=0; i<50; i++ {
		Ant_ID := i
		a, b := Ant()
		dis_ants = append(dis_ants, math.Round(a/100))
		map_distance[i] = math.Round(a/100)
		map_path[i] = b
		slice_ants_arrived = append(slice_ants_arrived, Ant_ID)
		fmt.Println("Le mapping est le suivant :", map_distance)
		for k:=0;k<len(slice_ants_arrived);k++{
			map_distance[slice_ants_arrived[k]] = map_distance[slice_ants_arrived [k]]-1
			if (map_distance[slice_ants_arrived[k]]==0){
				fmt.Println("The ant N°", slice_ants_arrived[k], "just arrived")
				for m:=0;m<Nnode-2;m++ {
					from := map_path[slice_ants_arrived[0]][m]
					to := map_path[slice_ants_arrived[0]][m+1]
					passij[from][to] = passij[from][to] + 1
					pij[from][to] = float64(1)/(float64(1)+math.Exp(-float64(passij[from][to])/float64(30)))
					//il faut penser a suprimer de slice_ants_arried la fourmi arrivée
				}
				delete(map_distance, slice_ants_arrived[k])
				delete(map_path, slice_ants_arrived[k])
			}
		}
	}
	//fmt.Println("Disntace parcourue par l'ant :", dis_ants)
	//fmt.Println("Le mapping est le suivant :", map_path)


}

func Ant() (float64, []int){
	var d_ant float64
	var ant_path []int
	slice_node := make([]int, Nnode-1)
	Start_node := 0
	for i:=1; i<len(slice_node)+1; i++{
		slice_node[i-1] = i
	}
	Current_node := Start_node
	for i:=0; i<Nnode-1; i++{
		//fmt.Println("Tirage N°", i+1)
		var ant_next int
		var ant_ind int
		ant_p := float64(rand.Intn(1000))/1000
		var ant_thresh float64
		for j:=0; j<len(slice_node); j++{
			//ant_thresh = ant_thresh + float64(1)/float64(len(slice_node))*pij[Current_node][j]
			var proba_current_node []float64
			var proba_norm_current_node float64
			for k:=0;k<len(slice_node);k++{
				proba_norm_current_node = proba_norm_current_node + pij[Current_node][slice_node[k]]
			}
			for k:=0;k<len(slice_node);k++{
				proba_current_node = append(proba_current_node, float64(pij[Current_node][slice_node[k]])/float64(proba_norm_current_node))
			}
			ant_thresh = ant_thresh + proba_current_node[j]
			if ant_p <= ant_thresh {
				//fmt.Println("Le reste des noeuds :", slice_node)
				//fmt.Println("La proba que la fourmis tire :", ant_p)
				//fmt.Println("Le thresh de la fourmi :", ant_thresh)
				//fmt.Println("L'indice du noeud choisi:", j)
				ant_ind = j
				ant_next = slice_node[j]
				ant_path = append(ant_path, ant_next)
				//fmt.Println("Le chemin parcouru est:",ant_path)
				break
			}
		}
		//fmt.Println("L'indice ant est :", ant_ind, "\n\n")
		if ant_ind == len(slice_node){
			slice_node = slice_node[:ant_ind]
		}
		if ant_ind == 0{
			slice_node = slice_node[(ant_ind+1):]
		}
		if ant_ind != 0 && ant_ind != len(slice_node) {
			slice_node = append(slice_node[:ant_ind], slice_node[(ant_ind+1):]...)
		}


	}
	for i:=0;i<len(ant_path)-1;i++{
		d_ant = d_ant + dis[ant_path[i]][ant_path[i+1]]
	}
	return d_ant, ant_path
}

