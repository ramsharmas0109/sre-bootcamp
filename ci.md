1. Code
2. Build Binary
3. Test DB up
4. Test DB migrate 
5. Run Tests 
6. Docker build
7. Docker tag
8. Docker push
9. Scan Images

Continious integration is all about this, that we automate the above 9 steps so that we can trigger it from anywhere, without having any dependency on a developer or machine. 

-> Reproducibility -> We can release our application -> docker tag with latest changes in more predicability.
-> Dependency on devs & specfic machines goes down.
-> Reduction in human error, as automation will be incharge of running these steps.
-> Security improves as we catch bugs using tests & scans. 
-> Less manual work.

Abstraction: High level command -> Understand -> Work

When these steps should get triggered?
- events ?? such that I trigger my pipelines. 

Pipelines has steps/ aka stages, and we look for certain events if that event happens then we trigger the pipeline.

Events? 

Whenever there is any change in the code, then the pipeline should get triggered.

A -> Valid
B -> Valid
C -> Not valid

Whenever there is any change in the code, then the pipeline should get triggered. This is not 100% valid all the time, it depends on the validity of the changes.

Every stage/ step in the pipeline depends on the success of the previous stage. As in if binary is not building, then running tests wont make sense. 

Similarly if Run tests are failing then buidling docker images wont make sense. 


If two engg are working on a project then what will be the event?


A -> changes made and pushed this qualifies as event
B -> event 


main (prodcution) -> client 

A's laptop
feature-A from main
code
test
validate changes
PR

Changes pushed to my feature branch can qualify as an event


main (prodcution) -> client 

B's laptop
feature-B from main
code
test
validate changes
PR

Changes pushed to my feature branch can qualify as an event


Git
main (prod) -> Client
feature-a -> A
feature-b -> B

Let's merge both A and B's PR in main. Main commit, push -> Push event -> Trigger

Events:
Push to feature branch
Merge to main

Pipeline that has some steps that gets trigger on two events (push on feature branch and merge on main)
Pipeline will have some steps that will run, and we need to have some kind of checks so that next step only runs if the
previous one is successful

If my pipleine is successful then someone should get informed, also if my pipline fails again someone should get informed.

Tools 
CicrleCI, TravisCI, GitHub Actions, Jenkins, Gitlab CI, Concourse CI... 

- Pipeline -> DONE
- Event -> DONE
- Stages -> DONE

Where should this pipeline run? 

Github Actions


YAML -> Declarative

Declarative vs Imeperative

