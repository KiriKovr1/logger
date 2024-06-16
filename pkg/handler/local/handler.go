package local

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

type LocalHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type LocalHandler struct {
	slog.Handler
	l     *log.Logger
	attrs []slog.Attr
}

func (local *LocalHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range local.attrs {
		fields[a.Key] = a.Value.Any()
	}

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	local.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func (local *LocalHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LocalHandler{
		Handler: local.Handler,
		l:       local.l,
		attrs:   attrs,
	}
}

func (h *LocalHandler) WithGroup(name string) slog.Handler {
	return &LocalHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}

func New(out io.Writer, opts LocalHandlerOptions) *LocalHandler {
	return &LocalHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}
}
