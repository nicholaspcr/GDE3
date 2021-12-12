from pymoo.factory import get_problem, get_reference_directions, get_visualization
from pymoo.util.plotting import plot


x = [0.040971105531507235,0.550373235584878,0.6817311625009819,0.6274478938025135,0.9234111071427142,0.02499901960750534,0.136171616578574,0.9084459589232222,0.21089363254881652,0.08574450529306678,0.20551052286248087,0.43442188671029464]
n_var = len(x)
n_obj = 3

problems = ["dtlz1", "dtlz2", "dtlz3", "dtlz4", "dtlz5", "dtlz6", "dtlz7"]

for problem in problems:
    p = get_problem(problem, n_var, n_obj)
    r = p.evaluate(x)
    print(r)



# PYMOO
# [  9.86928923   8.06270417 419.74214702]
# [1.35981478 1.59402999 0.13503034]
# [566.92680262 664.57456923  56.29613817]
# [2.09958592e+000 3.83693240e-026 5.83474399e-139]
# [1.41890847 1.54166358 0.13503034]
# [6.42107287 7.40537588 0.63167125]
# [ 0.04097111  0.55037324 17.32742807]

# GOLANG
# [9.869289225575503 8.062704169938133 419.74214702336616]
# [1.3598147826944689 1.5940299863405385 0.13503034348631712]
# [566.9268026207471 664.5745692269643 56.296138168016384]
# [2.0995859197111355 3.8369323955770535e-26 5.834743988703213e-139]
# [1.4189084710787399 1.5416635791534465 0.13503034348631712]
# [6.421072865759349 7.405375882273675 0.6316712455606305]
# [0.040971105531507235 0.550373235584878 17.32742807181844]
