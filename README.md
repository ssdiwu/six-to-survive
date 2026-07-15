# Six to Survive · 撑到六点

> 飞书消息从四面八方涌来，你在工位上疯狂敲键盘清活、甩锅、摸鱼，撑到 18:00 下班就算赢。

## 是什么

浏览器 2D 弹幕 Roguelike。玩家固定在屏幕中央工位圈（WASD 小范围移动），按任意键打出字符射向最近的飞书消息，空格摸鱼回血，Shift 甩锅反弹全屏弹幕。时间从 9:00 走到 18:00，撑到下班胜利。

本项目用于豆包 **Seed-Evolving** 模型开发者测评：在 TRAE 里让 Evolving 一次成型交付完整可玩的网页游戏，体现生产级 Coding 能力。

## 当前状态

当前处于 **v0.1 规格基线**：产品玩法和工程约束已经写入 PRD，游戏主体尚未实现。`index.html` 目前只是可打开的 Canvas 2D 骨架，用于承接下一步最小可玩切片。

## 怎么跑

无构建、无必需外部依赖。直接用浏览器打开 `index.html` 即可，或起个静态服务：

```bash
# Python
python3 -m http.server 8000
# 然后打开 http://localhost:8000
```

## 目录

```
six-to-survive/
├── index.html          # 单文件入口（最终交付）
├── scripts/            # 开发期按需创建的 JS 模块，最终合入 index.html
├── styles/             # 开发期按需创建的样式，最终合入 index.html
├── assets/             # 开发期按需创建的素材，程序化生成优先
├── doc/                # 设计文档与 PRD
│   ├── README.md
│   └── 00-产品与原则/
│       └── PRD.md
├── AGENTS.md           # AI 协作规范
├── CHANGELOG.md        # 面向玩家的版本变更
├── .gitignore          # 本地文件忽略规则
└── README.md           # 本文件
```

## 开发原则

- **单文件交付优先**：最终版收敛到 `index.html`（HTML + CSS + JS 全内联），便于 TRAE/Evolving 一次成型和录屏
- **开发期可拆分**：`scripts/` `styles/` 仅用于开发调试，合入时内联
- **表驱动**：敌人、词条、波次全部用 JS 对象/数组描述，不靠硬编码
- **零必需外部依赖**：不引入 npm 包，不依赖外部图片、音频或字体才能运行；可选 CDN 字体只能作为增强并保留系统字体兜底，不使用 Tailwind
- **Canvas 2D**：不使用 three.js/PixiJS，保持一次成型率高

## 操控

| 键 | 行为 |
|---|---|
| 任意字母/数字键 | 打出一个字符，射向最近的飞书敌人 |
| WASD / 方向键 | 在工位圈内小范围移动 |
| 空格（长按） | 摸鱼：隐身回 SAN/精力，时间停走，KPI 掉 |
| Shift | 甩锅：全屏弹幕反弹为友方穿透弹，3 秒不能攻击 |

## 文档

详见 `doc/README.md`。
