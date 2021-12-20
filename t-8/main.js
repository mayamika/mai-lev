paper.install(window);

$(function () {
    if (!('getContext' in document.createElement('canvas'))) {
        alert('Sorry, it looks like your browser does not support canvas!');
        return false;
    }

    const canvas = $('canvas')[0];
    paper.setup(canvas);

    const length = 15;
    let points = 5;

    const doc = $(document)
    var tool = new Tool();

    let initialized = false;
    let myPath;

    function createWorm(color) {
        const path = new paper.Path({
            strokeColor: color,
            strokeWidth: 20,
            strokeCap: 'round'
        });

        const start = new paper.Point(Math.random() * 100, Math.random() * 100);
        for (var i = 0; i < points; i++) {
            path.add(new paper.Point(i * length + start.x, 0 + start.y));
        }

        return path;
    }

    function randomColor() {
        const colors = ['#5C4B51', '#8CBEB2', '#F3B562', '#F06060']
        return colors[Math.floor(Math.random() * colors.length)];
    }

    var foods = [];
    const spawnFood = () => {
        const food = new Path.Circle({
            center: [Math.random() * 1500, Math.random() * 800],
            radius: 10,
            fillColor: randomColor()
        });

        foods.push(food);
    };

    let weightToFill = 0;
    const eatFood = () => {
        const head = myPath.firstSegment.point;

        for (const [i, food] of foods.entries()) {
            if (food.strokeBounds.contains(head)) {
                myPath.strokeColor = food.fillColor;

                foods.splice(i, 1);
                food.remove();

                weightToFill++;
            }
        }
    };

    paper.tool.onMouseMove = function (event) {
        if (!initialized) {
            const color = randomColor();

            myPath = createWorm(color);
            initialized = true;
        }

        const tail = myPath.lastSegment.point;

        myPath.firstSegment.point = event.point;
        for (var i = 0; i < points - 1; i++) {
            var segment = myPath.segments[i];
            var nextSegment = segment.next;
            var vector = new paper.Point(segment.point.x - nextSegment.point.x, segment.point.y - nextSegment.point.y);
            vector.length = length;
            nextSegment.point = new paper.Point(segment.point.x - vector.x, segment.point.y - vector.y);
        }
        myPath.smooth();

        eatFood();
        if (weightToFill > 0) {
            myPath.add(tail);
            --weightToFill;
            ++points;
        }
    }

    setInterval(spawnFood, 500);

    function tick() {
        paper.view.draw();
        requestAnimationFrame(tick);
    }
    requestAnimationFrame(tick);
});
