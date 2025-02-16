package impl

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/kubeagi/arcadia/apiserver/graph/generated"
	"github.com/kubeagi/arcadia/apiserver/pkg/application"
	"github.com/kubeagi/arcadia/apiserver/pkg/llm"
	"github.com/kubeagi/arcadia/apiserver/pkg/rag"
)

// Rag is the resolver for the RAG field.
func (r *mutationResolver) Rag(ctx context.Context) (*generated.RAGMutation, error) {
	return &generated.RAGMutation{}, nil
}

// Rag is the resolver for the RAG field.
func (r *queryResolver) Rag(ctx context.Context) (*generated.RAGQuery, error) {
	return &generated.RAGQuery{}, nil
}

// Application is the resolver for the application field.
func (r *rAGResolver) Application(ctx context.Context, obj *generated.Rag) (*generated.Application, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	r1, err := rag.GetV1alpha1RAG(ctx, c, obj.Name, obj.Namespace)
	if err != nil {
		return nil, err
	}
	ns := obj.Namespace
	if r1.Spec.Application.Namespace != nil {
		ns = *r1.Spec.Application.Namespace
	}
	return application.GetApplication(ctx, c, r1.Spec.Application.Name, ns)
}

// Datasets is the resolver for the datasets field.
func (r *rAGResolver) Datasets(ctx context.Context, obj *generated.Rag) ([]*generated.RAGDataset, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return rag.GetRAGDatasets(ctx, c, obj.Name, obj.Namespace)
}

// JudgeLlm is the resolver for the judgeLLM field.
func (r *rAGResolver) JudgeLlm(ctx context.Context, obj *generated.Rag) (*generated.Llm, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	r1, err := rag.GetV1alpha1RAG(ctx, c, obj.Name, obj.Namespace)
	if err != nil {
		return nil, err
	}
	ns := obj.Namespace
	if r1.Spec.Application.Namespace != nil {
		ns = *r1.Spec.Application.Namespace
	}
	return llm.ReadLLM(ctx, c, r1.Spec.JudgeLLM.Name, ns)
}

// Metrics is the resolver for the metrics field.
func (r *rAGResolver) Metrics(ctx context.Context, obj *generated.Rag) ([]*generated.RAGMetric, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return rag.GetRAGMetrics(ctx, c, obj.Name, obj.Namespace)
}

// CreateRag is the resolver for the createRAG field.
func (r *rAGMutationResolver) CreateRag(ctx context.Context, obj *generated.RAGMutation, input generated.CreateRAGInput) (*generated.Rag, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return rag.CreateRAG(ctx, c, &input)
}

// UpdateRag is the resolver for the updateRAG field.
func (r *rAGMutationResolver) UpdateRag(ctx context.Context, obj *generated.RAGMutation, input generated.UpdateRAGInput) (*generated.Rag, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return rag.UpdateRAG(ctx, c, &input)
}

// DeleteRag is the resolver for the deleteRAG field.
func (r *rAGMutationResolver) DeleteRag(ctx context.Context, obj *generated.RAGMutation, input generated.DeleteRAGInput) (*string, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return nil, rag.DeleteRAG(ctx, c, &input)
}

// GetRag is the resolver for the getRAG field.
func (r *rAGQueryResolver) GetRag(ctx context.Context, obj *generated.RAGQuery, name string, namespace string) (*generated.Rag, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return rag.GetRAG(ctx, c, name, namespace)
}

// ListRag is the resolver for the listRAG field.
func (r *rAGQueryResolver) ListRag(ctx context.Context, obj *generated.RAGQuery, input generated.ListRAGInput) (*generated.PaginatedResult, error) {
	c, err := getClientFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	return rag.ListRAG(ctx, c, &input)
}

// RAG returns generated.RAGResolver implementation.
func (r *Resolver) RAG() generated.RAGResolver { return &rAGResolver{r} }

// RAGMutation returns generated.RAGMutationResolver implementation.
func (r *Resolver) RAGMutation() generated.RAGMutationResolver { return &rAGMutationResolver{r} }

// RAGQuery returns generated.RAGQueryResolver implementation.
func (r *Resolver) RAGQuery() generated.RAGQueryResolver { return &rAGQueryResolver{r} }

type rAGResolver struct{ *Resolver }
type rAGMutationResolver struct{ *Resolver }
type rAGQueryResolver struct{ *Resolver }
