package chant

import(
	"time"
	"math/rand"
	"strconv"
)

type PSOParam struct{
	Large			int			//地图大小
	ParticleNum		int			//粒子数量
	Matrix			int			//探索矩阵大小
	Step 			int			//线性粒子间距
}

type Particle struct{
	 name		string
	 index		int
}		//粒子
var(
	Map  []int
	i	int
	Numcount = make(chan int)
)
	

func (Pso *PSOParam)CreateMap()([]int){		//创建地图
	
	rand.Seed(time.Now().Unix())
	for i = 0;i < Pso.Large; i++{
		Map = append(Map,rand.Intn(Pso.Large))
	}
	return Map 

}

func (Pso *PSOParam)CreateParticle()([]Particle){		//创建粒子群
	var(
		i 		  			int
		ParticleSwarm  	=   make([]Particle,Pso.ParticleNum)
	 )
	 for i = 1 ; i <= Pso.ParticleNum; i++{
		ParticleSwarm[i-1].name 	= "part"+strconv.Itoa(i)
		if i == 1{
			ParticleSwarm[i-1].index	= Pso.Step 
		}
		ParticleSwarm[i-1].index	= Pso.Step * (i-1)
	 }
	  
	return  ParticleSwarm
}

func (Pso *PSOParam)SearchNum(Smap []int,Sparticle *[]Particle,targetNum int) (int,int){
			for key , particleNum := range *Sparticle{
				
		 		for i = particleNum.index; i< particleNum.index+Pso.Step;i++{
						if Smap[i] == targetNum {
							return Smap[i],i
						}
				}
				if key+Pso.Step > len(*Sparticle){
					for i= len(Smap)-1; i> len(Smap)-1-Pso.Step;i--{
						if Smap[i] == targetNum {
							return Smap[i],i
						}
					}
				}
			}
		return -1,0
}

