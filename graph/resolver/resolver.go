package resolver

import (
	"TaskManager/graph/generated"
	"TaskManager/graph/model"
	"TaskManager/models"
	"context"

	"github.com/jinzhu/gorm"
)

type Resolver struct {
    DB *gorm.DB
}


func (r *Resolver) Mutation()  generated.MutationResolver{
    return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
    return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
    var tasks []*models.Task
    r.DB.Find(&tasks)
    var result []*model.Task
    for _, task := range tasks {
        result = append(result, &model.Task{
            ID:          task.ID,
            Title:       task.Title,
            Description: task.Description,
            Completed:   task.Completed,
        })
    }
    return result, nil
}

func (r *mutationResolver) CreateTask(ctx context.Context, title string, description string) (*model.Task, error) {
    task := &models.Task{Title: title, Description: description}
    r.DB.Create(task)
    return &model.Task{
        ID:          task.ID,
        Title:       task.Title,
        Description: task.Description,
        Completed:   task.Completed,
    }, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, completed bool) (*model.Task, error) {
    var task models.Task
    r.DB.First(&task, id)
    task.Completed = completed
    r.DB.Save(&task)
    return &model.Task{
        ID:          task.ID,
        Title:       task.Title,
        Description: task.Description,
        Completed:   task.Completed,
    }, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (bool, error) {
    var task models.Task
    r.DB.First(&task, id)
    r.DB.Delete(&task)
    return true, nil
}
